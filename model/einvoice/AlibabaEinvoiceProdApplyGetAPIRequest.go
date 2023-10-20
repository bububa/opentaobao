package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceprodapplygetAPIRequest 查询发票申请 API请求
// alibaba.einvoice.prod.apply.get
//
// 查询申请的详细信息，包含申请所关联的发票摘要信息+板式文件+预览图；
// 场景使用：1、业务前台收到申请状态变更消息后，调用此接口查询申请详情；2、主动补偿查询：当指定了自动开票，且发票申请长时间未收到状态变更通知时，可能存在丢消息的情况，此时可主动查询该申请，然后更新本地工单状态。
type AlibabaeinvoiceprodapplygetAPIRequest struct {
	model.Params
	// 查询申请请求
	_invoiceApplyQueryDto *InvoiceApplyDtlQueryDto
}

// NewAlibabaeinvoiceprodapplygetRequest 初始化AlibabaeinvoiceprodapplygetAPIRequest对象
func NewAlibabaeinvoiceprodapplygetRequest() *AlibabaeinvoiceprodapplygetAPIRequest {
	return &AlibabaeinvoiceprodapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceprodapplygetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.prod.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceprodapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceprodapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceApplyQueryDto is InvoiceApplyQueryDto Setter
// 查询申请请求
func (r *AlibabaeinvoiceprodapplygetAPIRequest) SetInvoiceApplyQueryDto(_invoiceApplyQueryDto *InvoiceApplyDtlQueryDto) error {
	r._invoiceApplyQueryDto = _invoiceApplyQueryDto
	r.Set("invoice_apply_query_dto", _invoiceApplyQueryDto)
	return nil
}

// GetInvoiceApplyQueryDto InvoiceApplyQueryDto Getter
func (r AlibabaeinvoiceprodapplygetAPIRequest) GetInvoiceApplyQueryDto() *InvoiceApplyDtlQueryDto {
	return r._invoiceApplyQueryDto
}
