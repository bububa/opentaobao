package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceFlowTaxCreate 创建税控开通工单
// alibaba.einvoice.flow.tax.create
//
// 商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。
func AlibabaEinvoiceFlowTaxCreate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceFlowTaxCreateAPIRequest, resp *einvoice.AlibabaEinvoiceFlowTaxCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
