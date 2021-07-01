package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkOrderRefundList
五道口交易退款批量查询
alibaba.wdk.order.refund.list

按照条件查询退款数据 */
func AlibabaWdkOrderRefundList(clt *core.SDKClient, req *wdk.AlibabaWdkOrderRefundListAPIRequest, session string) (*wdk.AlibabaWdkOrderRefundListAPIResponse, error) {
	var resp wdk.AlibabaWdkOrderRefundListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
