package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

/* TaobaoAlitripSellerModifyList
【机票代理商订单】改签订单列表
taobao.alitrip.seller.modify.list

提供机票代理商查询改签订单列表 */
func TaobaoAlitripSellerModifyList(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerModifyListAPIRequest, session string) (*jipiao.TaobaoAlitripSellerModifyListAPIResponse, error) {
	var resp jipiao.TaobaoAlitripSellerModifyListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
