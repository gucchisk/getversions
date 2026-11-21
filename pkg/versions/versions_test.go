package versions_test

import (
	"testing"

	"github.com/gucchisk/getversions/pkg/versions"
	"github.com/stretchr/testify/assert"
)

func TestNewVersions(t *testing.T) {
	testCases := []struct {
		name             string
		inputVersions    []string
		expectedVersions []string
	}{
		{
			name:             "空のスライスの場合は空を返す",
			inputVersions:    []string{},
			expectedVersions: []string{},
		},
		{
			name:             "vプレフィックスの有無にかかわらず元の表記のまま保持される",
			inputVersions:    []string{"1.0.0", "2.3", "v1.2.3"},
			expectedVersions: []string{"1.0.0", "2.3", "v1.2.3"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := versions.NewVersions(testCase.inputVersions)

			// 初期状態では条件なし・最新のみではないため、全バージョンが返る
			assert.Equal(t, testCase.expectedVersions, actual.Get())
		})
	}
}

func TestFilter(t *testing.T) {
	inputVersions := []string{"1.0.0", "1.2.0", "v1.2.3", "2.0.0"}

	t.Run("同じインスタンスを返す", func(t *testing.T) {
		target := versions.NewVersions(inputVersions)

		assert.Same(t, target, target.Filter("1.2"))
		assert.Same(t, target, target.Filter(""))
	})

	t.Run("空の条件ではフィルタされない", func(t *testing.T) {
		actual := versions.NewVersions(inputVersions).Filter("").Get()

		assert.Equal(t, inputVersions, actual)
	})

	t.Run("条件のvプレフィックスの有無は結果に影響しない", func(t *testing.T) {
		withoutPrefix := versions.NewVersions(inputVersions).Filter("1.2").Get()
		withPrefix := versions.NewVersions(inputVersions).Filter("v1.2").Get()

		assert.Equal(t, []string{"1.2.0", "v1.2.3"}, withoutPrefix)
		assert.Equal(t, withoutPrefix, withPrefix)
	})
}

func TestOnlyLatest(t *testing.T) {
	inputVersions := []string{"1.0.0", "v2.0.0", "1.5.0"}

	t.Run("同じインスタンスを返す", func(t *testing.T) {
		target := versions.NewVersions(inputVersions)

		assert.Same(t, target, target.OnlyLatest())
	})

	t.Run("最新バージョンを元の表記のまま1件だけ返す", func(t *testing.T) {
		actual := versions.NewVersions(inputVersions).OnlyLatest().Get()

		assert.Equal(t, []string{"v2.0.0"}, actual)
	})
}

func TestGet(t *testing.T) {
	inputVersions := []string{"1.0.0", "1.2.0", "1.2.3", "1.20.0", "v2.0.0", "2.1.0"}

	testCases := []struct {
		name             string
		inputVersions    []string
		condition        string
		onlyLatest       bool
		expectedVersions []string
	}{
		{
			name:             "条件なしの場合は全バージョンを元の表記のまま返す",
			inputVersions:    inputVersions,
			expectedVersions: []string{"1.0.0", "1.2.0", "1.2.3", "1.20.0", "v2.0.0", "2.1.0"},
		},
		{
			name:             "条件がマイナーバージョンまでならマイナーバージョンまで一致するバージョンのみ返す",
			inputVersions:    inputVersions,
			condition:        "1.2",
			expectedVersions: []string{"1.2.0", "1.2.3"},
		},
		{
			name:             "条件がパッチバージョンまでなら完全一致するバージョンのみ返す",
			inputVersions:    inputVersions,
			condition:        "1.2.3",
			expectedVersions: []string{"1.2.3"},
		},
		{
			name:             "最新のみでマイナーバージョンまでの条件では前方一致する別マイナーバージョンを選ばない",
			inputVersions:    inputVersions,
			condition:        "1.2",
			onlyLatest:       true,
			expectedVersions: []string{"1.2.3"},
		},
		{
			name:             "vプレフィックス付きの条件でもフィルタできる",
			inputVersions:    inputVersions,
			condition:        "v2",
			expectedVersions: []string{"v2.0.0", "2.1.0"},
		},
		{
			name:             "条件に一致するバージョンがない場合は空を返す",
			inputVersions:    inputVersions,
			condition:        "3",
			expectedVersions: []string{},
		},
		{
			name:             "最新のみの場合は最大バージョンを返す",
			inputVersions:    inputVersions,
			onlyLatest:       true,
			expectedVersions: []string{"2.1.0"},
		},
		{
			name:             "最新のみの場合は入力順に関係なく最大バージョンを返す",
			inputVersions:    []string{"1.10.0", "1.9.9", "1.2.0"},
			onlyLatest:       true,
			expectedVersions: []string{"1.10.0"},
		},
		{
			name:             "条件と最新のみを組み合わせると条件内の最大バージョンを返す",
			inputVersions:    inputVersions,
			condition:        "1",
			onlyLatest:       true,
			expectedVersions: []string{"1.20.0"},
		},
		{
			name:             "最新のみで条件に一致するバージョンがない場合は空を返す",
			inputVersions:    inputVersions,
			condition:        "3",
			onlyLatest:       true,
			expectedVersions: []string{},
		},
		{
			name:             "最新のみでバージョンが空の場合は空を返す",
			inputVersions:    []string{},
			onlyLatest:       true,
			expectedVersions: []string{},
		},
		{
			name:             "最新のみでsemverとして不正なバージョンは無視される",
			inputVersions:    []string{"invalid", "1.0.0"},
			onlyLatest:       true,
			expectedVersions: []string{"1.0.0"},
		},
		{
			name:             "最新のみでv0.0.0しかない場合は空を返す",
			inputVersions:    []string{"0.0.0"},
			onlyLatest:       true,
			expectedVersions: []string{},
		},
		{
			name:             "条件なしでバージョンが空の場合は空を返す",
			inputVersions:    []string{},
			expectedVersions: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			target := versions.NewVersions(testCase.inputVersions).Filter(testCase.condition)
			if testCase.onlyLatest {
				target = target.OnlyLatest()
			}

			assert.Equal(t, testCase.expectedVersions, target.Get())
		})
	}
}
