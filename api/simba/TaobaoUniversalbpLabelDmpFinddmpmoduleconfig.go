package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpLabelDmpFinddmpmoduleconfig 查询dmp浮层配置
// taobao.universalbp.label.dmp.finddmpmoduleconfig
//
// 入参账号信息，出参达摩盘相关配置信息
func TaobaoUniversalbpLabelDmpFinddmpmoduleconfig(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest, resp *simba.TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
