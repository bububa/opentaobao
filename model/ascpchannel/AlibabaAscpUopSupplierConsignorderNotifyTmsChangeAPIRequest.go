package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest 商家修改运单号 API请求
// alibaba.ascp.uop.supplier.consignorder.notify.tms.change
//
// 供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改
type AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest struct {
	model.Params
	// 修改运单号请求模型
	_modifyMailNoRequest *Modifymailnorequest
}

// NewAlibabaascpuopsupplierconsignordernotifytmschangeRequest 初始化AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest对象
func NewAlibabaascpuopsupplierconsignordernotifytmschangeRequest() *AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest {
	return &AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.notify.tms.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyMailNoRequest is ModifyMailNoRequest Setter
// 修改运单号请求模型
func (r *AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest) SetModifyMailNoRequest(_modifyMailNoRequest *Modifymailnorequest) error {
	r._modifyMailNoRequest = _modifyMailNoRequest
	r.Set("modify_mail_no_request", _modifyMailNoRequest)
	return nil
}

// GetModifyMailNoRequest ModifyMailNoRequest Getter
func (r AlibabaascpuopsupplierconsignordernotifytmschangeAPIRequest) GetModifyMailNoRequest() *Modifymailnorequest {
	return r._modifyMailNoRequest
}
