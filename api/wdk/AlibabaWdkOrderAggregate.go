package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkOrderAggregate
淘鲜达订单按门店机台号聚合查询
alibaba.wdk.order.aggregate

淘鲜达订单按门店机台号聚合查询 */
func AlibabaWdkOrderAggregate(clt *core.SDKClient, req *wdk.AlibabaWdkOrderAggregateAPIRequest, session string) (*wdk.AlibabaWdkOrderAggregateAPIResponse, error) {
	var resp wdk.AlibabaWdkOrderAggregateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
