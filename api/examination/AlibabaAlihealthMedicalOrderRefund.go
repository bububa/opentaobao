package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthmedicalorderrefund 退款接口
// alibaba.alihealth.medical.order.refund
//
// 退款接口
func Alibabaalihealthmedicalorderrefund(clt *core.SDKClient, req *examination.AlibabaalihealthmedicalorderrefundAPIRequest, session string) (*examination.AlibabaalihealthmedicalorderrefundAPIResponse, error) {
	var resp examination.AlibabaalihealthmedicalorderrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
