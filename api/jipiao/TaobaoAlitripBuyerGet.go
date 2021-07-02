package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripBuyerGet 敏感信息查询
// taobao.alitrip.buyer.get
//
// 针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
func TaobaoAlitripBuyerGet(clt *core.SDKClient, req *jipiao.TaobaoAlitripBuyerGetAPIRequest, session string) (*jipiao.TaobaoAlitripBuyerGetAPIResponse, error) {
	var resp jipiao.TaobaoAlitripBuyerGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
