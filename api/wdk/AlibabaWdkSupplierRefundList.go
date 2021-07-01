package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkSupplierRefundList
五道口按供应商拉取退款单
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单 */
func AlibabaWdkSupplierRefundList(clt *core.SDKClient, req *wdk.AlibabaWdkSupplierRefundListAPIRequest, session string) (*wdk.AlibabaWdkSupplierRefundListAPIResponse, error) {
	var resp wdk.AlibabaWdkSupplierRefundListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
