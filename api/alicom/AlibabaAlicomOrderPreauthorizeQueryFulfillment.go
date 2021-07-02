package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomOrderPreauthorizeQueryFulfillment 履约结果查询
// alibaba.alicom.order.preauthorize.query.fulfillment
//
// 预授权-履约结果查询
func AlibabaAlicomOrderPreauthorizeQueryFulfillment(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest, session string) (*alicom.AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIResponse, error) {
	var resp alicom.AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
