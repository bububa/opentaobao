package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest 预售商家仓出库 API请求
// alibaba.ascp.uop.taobao.presalesorder.consignconfirm
//
// 预售商家仓出库
type AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest struct {
	model.Params
	// 预售订单商家仓出库对象
	_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest
}

// NewAlibabaascpuoptaobaopresalesorderconsignconfirmRequest 初始化AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest对象
func NewAlibabaascpuoptaobaopresalesorderconsignconfirmRequest() *AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest {
	return &AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.taobao.presalesorder.consignconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPresalesOrderConsignConfirmRequest is PresalesOrderConsignConfirmRequest Setter
// 预售订单商家仓出库对象
func (r *AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest) SetPresalesOrderConsignConfirmRequest(_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest) error {
	r._presalesOrderConsignConfirmRequest = _presalesOrderConsignConfirmRequest
	r.Set("presales_order_consign_confirm_request", _presalesOrderConsignConfirmRequest)
	return nil
}

// GetPresalesOrderConsignConfirmRequest PresalesOrderConsignConfirmRequest Getter
func (r AlibabaascpuoptaobaopresalesorderconsignconfirmAPIRequest) GetPresalesOrderConsignConfirmRequest() *Presalesorderconsignconfirmrequest {
	return r._presalesOrderConsignConfirmRequest
}
