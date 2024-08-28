package jipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerModifyList 【机票代理商订单】改签订单列表
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
func TaobaoAlitripSellerModifyList(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerModifyListAPIRequest, resp *jipiao.TaobaoAlitripSellerModifyListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
