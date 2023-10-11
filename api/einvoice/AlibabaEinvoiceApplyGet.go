package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceApplyGet 开票申请数据获取接口
// alibaba.einvoice.apply.get
//
// ERP获取开票申请数据
func AlibabaEinvoiceApplyGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceApplyGetAPIRequest, session string) (*einvoice.AlibabaEinvoiceApplyGetAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceApplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
