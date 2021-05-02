package all

import (
	"github.com/v2fly/v2ray-core/v4/main/commands/all/api"
	"github.com/v2fly/v2ray-core/v4/main/commands/all/tls"
	"github.com/v2fly/v2ray-core/v4/main/commands/base"
)

// go:generate go run github.com/v2fly/v2ray-core/v4/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		cmdLove,
		tls.CmdTLS,
		cmdUUID,
		cmdVerify,

		// documents
		docFormat,
		docMerge,
	)
}
