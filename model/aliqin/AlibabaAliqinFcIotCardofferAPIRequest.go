package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotCardofferAPIRequest 查询物联网卡上订购的offer API请求
// alibaba.aliqin.fc.iot.cardoffer
//
// 查询物联网卡上订购的offer
type AlibabaAliqinFcIotCardofferAPIRequest struct {
	model.Params
	// 具体ICCID的值
	_billreal string
	// ICCID
	_billsource string
}

// NewAlibabaAliqinFcIotCardofferRequest 初始化AlibabaAliqinFcIotCardofferAPIRequest对象
func NewAlibabaAliqinFcIotCardofferRequest() *AlibabaAliqinFcIotCardofferAPIRequest {
	return &AlibabaAliqinFcIotCardofferAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.cardoffer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillreal is Billreal Setter
// 具体ICCID的值
func (r *AlibabaAliqinFcIotCardofferAPIRequest) SetBillreal(_billreal string) error {
	r._billreal = _billreal
	r.Set("billreal", _billreal)
	return nil
}

// GetBillreal Billreal Getter
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetBillreal() string {
	return r._billreal
}

// SetBillsource is Billsource Setter
// ICCID
func (r *AlibabaAliqinFcIotCardofferAPIRequest) SetBillsource(_billsource string) error {
	r._billsource = _billsource
	r.Set("billsource", _billsource)
	return nil
}

// GetBillsource Billsource Getter
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetBillsource() string {
	return r._billsource
}
