package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellerrefundsearch 【机票代理商】退票申请单列表
// taobao.alitrip.seller.refund.search
//
// 查询退票申请单列表
func Taobaoalitripsellerrefundsearch(clt *core.SDKClient, req *jipiao.TaobaoalitripsellerrefundsearchAPIRequest, session string) (*jipiao.TaobaoalitripsellerrefundsearchAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellerrefundsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
