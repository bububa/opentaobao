package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellerrefundorderlistfetch 【机票代理商】——退票订单列表提取
// taobao.alitrip.seller.refundorderlist.fetch
//
// 代理商纬度退票订单列表提取
func Taobaoalitripsellerrefundorderlistfetch(clt *core.SDKClient, req *jipiao.TaobaoalitripsellerrefundorderlistfetchAPIRequest, session string) (*jipiao.TaobaoalitripsellerrefundorderlistfetchAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellerrefundorderlistfetchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
