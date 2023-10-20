package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensormakeup 美妆虚拟试装
// alibaba.interact.sensor.makeup
//
// 手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
func Alibabainteractsensormakeup(clt *core.SDKClient, req *interact.AlibabainteractsensormakeupAPIRequest, session string) (*interact.AlibabainteractsensormakeupAPIResponse, error) {
	var resp interact.AlibabainteractsensormakeupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
