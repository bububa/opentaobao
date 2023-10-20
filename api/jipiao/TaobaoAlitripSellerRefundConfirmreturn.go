package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellerrefundconfirmreturn 【机票代理商】确认退票
// taobao.alitrip.seller.refund.confirmreturn
//
// 确认退票
func Taobaoalitripsellerrefundconfirmreturn(clt *core.SDKClient, req *jipiao.TaobaoalitripsellerrefundconfirmreturnAPIRequest, session string) (*jipiao.TaobaoalitripsellerrefundconfirmreturnAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellerrefundconfirmreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
