// +build go1.8

package main

import (
	_ "gitlab.33.cn/chain33/bityuan/plugin"
	_ "gitlab.33.cn/chain33/chain33/plugin"
	_ "gitlab.33.cn/chain33/chain33/system"
	"gitlab.33.cn/chain33/chain33/util/cli"
)

func main() {
	cli.RunChain33("bityuan")
}
