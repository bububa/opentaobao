package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripbuyerget 敏感信息查询
// taobao.alitrip.buyer.get
//
// 针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
func Taobaoalitripbuyerget(clt *core.SDKClient, req *jipiao.TaobaoalitripbuyergetAPIRequest, session string) (*jipiao.TaobaoalitripbuyergetAPIResponse, error) {
	var resp jipiao.TaobaoalitripbuyergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
