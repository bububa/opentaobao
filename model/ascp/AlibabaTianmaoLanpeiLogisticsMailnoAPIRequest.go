package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest 阿里巴巴.天猫家装.揽配.物流.获取运单号 API请求
// alibaba.tianmao.lanpei.logistics.mailno
//
// 阿里巴巴.天猫家装.揽配.物流.获取运单号
type AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest struct {
	model.Params
	// scp单号
	_orderCode string
	// 包裹数量
	_packageQuantity int64
	// 货主ID
	_ownerId int64
}

// NewAlibabaTianmaoLanpeiLogisticsMailnoRequest 初始化AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest对象
func NewAlibabaTianmaoLanpeiLogisticsMailnoRequest() *AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest {
	return &AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.lanpei.logistics.mailno"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// scp单号
func (r *AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetPackageQuantity is PackageQuantity Setter
// 包裹数量
func (r *AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) SetPackageQuantity(_packageQuantity int64) error {
	r._packageQuantity = _packageQuantity
	r.Set("package_quantity", _packageQuantity)
	return nil
}

// GetPackageQuantity PackageQuantity Getter
func (r AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) GetPackageQuantity() int64 {
	return r._packageQuantity
}

// SetOwnerId is OwnerId Setter
// 货主ID
func (r *AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) SetOwnerId(_ownerId int64) error {
	r._ownerId = _ownerId
	r.Set("owner_id", _ownerId)
	return nil
}

// GetOwnerId OwnerId Getter
func (r AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest) GetOwnerId() int64 {
	return r._ownerId
}
