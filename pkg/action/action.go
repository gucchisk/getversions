package action

import (
	"bytes"
	"io"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/spf13/pflag"
)

type GetVersionsAction interface {
	Short() string
	Long() string
	GetVersions(reader io.Reader, flags pflag.FlagSet) []string
}

type GetVersionsPluginAction interface {
	Version() string
	Short() string
	Long() string
	FlagSet() pflag.FlagSet
	GetVersions(reader io.Reader, flags pflag.FlagSet) []string
}

type GetVersionsPluginActionRPC struct {
	client *rpc.Client
}

func (g *GetVersionsPluginActionRPC) Version() string {
	var resp string
	err := g.client.Call("Plugin.Version", struct{}{}, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

func (g *GetVersionsPluginActionRPC) Short() string {
	var resp string
	err := g.client.Call("Plugin.Short", struct{}{}, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

func (g *GetVersionsPluginActionRPC) Long() string {
	var resp string
	err := g.client.Call("Plugin.Long", struct{}{}, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

func (g *GetVersionsPluginActionRPC) FlagSet() pflag.FlagSet {
	var resp pflag.FlagSet
	err := g.client.Call("Plugin.FlagSet", struct{}{}, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type GetVersionsArgs struct {
	Data  []byte
	Flags pflag.FlagSet
}

func (g *GetVersionsPluginActionRPC) GetVersions(reader io.Reader, flags pflag.FlagSet) []string {
	var resp []string
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	args := GetVersionsArgs{
		Data:  data,
		Flags: flags,
	}
	err = g.client.Call("Plugin.GetVersions", args, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type Plugin struct {
	Impl GetVersionsPluginAction
}

func (s *Plugin) Version(args struct{}, resp *string) error {
	*resp = s.Impl.Version()
	return nil
}

func (s *Plugin) Short(args struct{}, resp *string) error {
	*resp = s.Impl.Short()
	return nil
}

func (s *Plugin) Long(args struct{}, resp *string) error {
	*resp = s.Impl.Long()
	return nil
}

func (s *Plugin) GetVersions(args GetVersionsArgs, resp *[]string) error {
	data := args.Data
	reader := bytes.NewReader(data)
	*resp = s.Impl.GetVersions(reader, args.Flags)
	return nil
}

type GetVersionsPlugin struct {
	Impl GetVersionsPluginAction
}

func (p *GetVersionsPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &Plugin{Impl: p.Impl}, nil
}

func (GetVersionsPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GetVersionsPluginActionRPC{client: c}, nil
}
