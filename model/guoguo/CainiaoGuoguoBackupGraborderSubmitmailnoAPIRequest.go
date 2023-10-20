package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguobackupgrabordersubmitmailnoAPIRequest 兜底派送订单的运单号回传接口 API请求
// cainiao.guoguo.backup.graborder.submitmailno
//
// 快递公司回传订单号和运单号给菜鸟裹裹
type CainiaoguoguobackupgrabordersubmitmailnoAPIRequest struct {
	model.Params
	// 菜鸟物流订单号
	_orderCode string
	// 运单号
	_mailNo string
}

// NewCainiaoguoguobackupgrabordersubmitmailnoRequest 初始化CainiaoguoguobackupgrabordersubmitmailnoAPIRequest对象
func NewCainiaoguoguobackupgrabordersubmitmailnoRequest() *CainiaoguoguobackupgrabordersubmitmailnoAPIRequest {
	return &CainiaoguoguobackupgrabordersubmitmailnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.backup.graborder.submitmailno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 菜鸟物流订单号
func (r *CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("orderCode", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mailNo", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoguoguobackupgrabordersubmitmailnoAPIRequest) GetMailNo() string {
	return r._mailNo
}
