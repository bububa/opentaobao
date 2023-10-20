package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractcurrentmixusernick 手淘混淆nick开放接口鉴权专用
// alibaba.interact.current.mixusernick
//
// 手淘混淆nick开放接口鉴权专用，无数据输入输出。
func Alibabainteractcurrentmixusernick(clt *core.SDKClient, req *interact.AlibabainteractcurrentmixusernickAPIRequest, session string) (*interact.AlibabainteractcurrentmixusernickAPIResponse, error) {
	var resp interact.AlibabainteractcurrentmixusernickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
