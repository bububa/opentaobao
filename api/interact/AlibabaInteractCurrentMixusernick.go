package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractCurrentMixusernick 手淘混淆nick开放接口鉴权专用
// alibaba.interact.current.mixusernick
//
// 手淘混淆nick开放接口鉴权专用，无数据输入输出。
func AlibabaInteractCurrentMixusernick(clt *core.SDKClient, req *interact.AlibabaInteractCurrentMixusernickAPIRequest, resp *interact.AlibabaInteractCurrentMixusernickAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
