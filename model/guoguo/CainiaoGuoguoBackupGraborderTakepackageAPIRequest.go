package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoBackupGraborderTakepackageAPIRequest 兜底派送订单的揽件接口 API请求
// cainiao.guoguo.backup.graborder.takepackage
//
// 快递公司回传订单号和四位取件码给菜鸟裹裹
type CainiaoGuoguoBackupGraborderTakepackageAPIRequest struct {
	model.Params
	// 物流订单号
	_orderCode string
	// 包裹四位码
	_packageCode string
}

// NewCainiaoGuoguoBackupGraborderTakepackageRequest 初始化CainiaoGuoguoBackupGraborderTakepackageAPIRequest对象
func NewCainiaoGuoguoBackupGraborderTakepackageRequest() *CainiaoGuoguoBackupGraborderTakepackageAPIRequest {
	return &CainiaoGuoguoBackupGraborderTakepackageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoBackupGraborderTakepackageAPIRequest) GetApiMethodName() string {
	return "cainiao.guoguo.backup.graborder.takepackage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoBackupGraborderTakepackageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGuoguoBackupGraborderTakepackageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 物流订单号
func (r *CainiaoGuoguoBackupGraborderTakepackageAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("orderCode", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoGuoguoBackupGraborderTakepackageAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetPackageCode is PackageCode Setter
// 包裹四位码
func (r *CainiaoGuoguoBackupGraborderTakepackageAPIRequest) SetPackageCode(_packageCode string) error {
	r._packageCode = _packageCode
	r.Set("packageCode", _packageCode)
	return nil
}

// GetPackageCode PackageCode Getter
func (r CainiaoGuoguoBackupGraborderTakepackageAPIRequest) GetPackageCode() string {
	return r._packageCode
}
