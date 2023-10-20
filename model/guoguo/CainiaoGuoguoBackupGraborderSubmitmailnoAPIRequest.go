package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest 兜底派送订单的运单号回传接口 API请求
// cainiao.guoguo.backup.graborder.submitmailno
//
// 快递公司回传订单号和运单号给菜鸟裹裹
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest struct {
	model.Params
	// 菜鸟物流订单号
	_orderCode string
	// 运单号
	_mailNo string
}

// NewCainiaoGuoguoBackupGraborderSubmitmailnoRequest 初始化CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest对象
func NewCainiaoGuoguoBackupGraborderSubmitmailnoRequest() *CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest {
	return &CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.backup.graborder.submitmailno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 菜鸟物流订单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("orderCode", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mailNo", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetMailNo() string {
	return r._mailNo
}
