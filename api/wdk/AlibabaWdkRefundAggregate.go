package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkRefundAggregate
淘鲜达退款单按门店聚合查询
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询 */
func AlibabaWdkRefundAggregate(clt *core.SDKClient, req *wdk.AlibabaWdkRefundAggregateAPIRequest, session string) (*wdk.AlibabaWdkRefundAggregateAPIResponse, error) {
	var resp wdk.AlibabaWdkRefundAggregateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
