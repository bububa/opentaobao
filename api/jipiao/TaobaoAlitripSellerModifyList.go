package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellermodifylist 【机票代理商订单】改签订单列表
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
func Taobaoalitripsellermodifylist(clt *core.SDKClient, req *jipiao.TaobaoalitripsellermodifylistAPIRequest, session string) (*jipiao.TaobaoalitripsellermodifylistAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellermodifylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
