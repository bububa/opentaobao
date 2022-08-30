package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthMedicalOrderRefund 退款接口
// alibaba.alihealth.medical.order.refund
//
// 退款接口
func AlibabaAlihealthMedicalOrderRefund(clt *core.SDKClient, req *examination.AlibabaAlihealthMedicalOrderRefundAPIRequest, session string) (*examination.AlibabaAlihealthMedicalOrderRefundAPIResponse, error) {
	var resp examination.AlibabaAlihealthMedicalOrderRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
