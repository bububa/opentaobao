package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellerrefundget 【机票代理商】退票申请单详情
// taobao.alitrip.seller.refund.get
//
// 查询退票申请单详情
func Taobaoalitripsellerrefundget(clt *core.SDKClient, req *jipiao.TaobaoalitripsellerrefundgetAPIRequest, session string) (*jipiao.TaobaoalitripsellerrefundgetAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellerrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
