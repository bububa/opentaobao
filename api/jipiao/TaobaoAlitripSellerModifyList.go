package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerModifyList 【机票代理商订单】改签订单列表
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
func TaobaoAlitripSellerModifyList(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerModifyListAPIRequest, resp *jipiao.TaobaoAlitripSellerModifyListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
