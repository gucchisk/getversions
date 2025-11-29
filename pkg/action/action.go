package action

import (
	"bytes"
	"io"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type GetVersionsAction interface {
	Short() string
	Long() string
	GetVersions(reader io.Reader) []string
}

type GetVersionsPluginAction interface {
	Version() string
	Short() string
	Long() string
	GetVersions(reader io.Reader) []string
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

func (g *GetVersionsPluginActionRPC) GetVersions(reader io.Reader) []string {
	var resp []string
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	err = g.client.Call("Plugin.GetVersions", data, &resp)
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

func (s *Plugin) GetVersions(data []byte, resp *[]string) error {
	reader := bytes.NewReader(data)
	*resp = s.Impl.GetVersions(reader)
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
