package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractshopfavor 店铺收藏
// alibaba.interact.shop.favor
//
// 店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
func Alibabainteractshopfavor(clt *core.SDKClient, req *interact.AlibabainteractshopfavorAPIRequest, session string) (*interact.AlibabainteractshopfavorAPIResponse, error) {
	var resp interact.AlibabainteractshopfavorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
