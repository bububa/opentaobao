package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackageentryorderconfirmAPIRequest 包裹入库单确认 API请求
// taobao.logistics.wms.packageentryorder.confirm
//
// 包裹入库单确认
type TaobaologisticswmspackageentryorderconfirmAPIRequest struct {
	model.Params
	// 请求
	_confirmPackageEntryOrderRequest *ConfirmPackageOrderRequest
}

// NewTaobaologisticswmspackageentryorderconfirmRequest 初始化TaobaologisticswmspackageentryorderconfirmAPIRequest对象
func NewTaobaologisticswmspackageentryorderconfirmRequest() *TaobaologisticswmspackageentryorderconfirmAPIRequest {
	return &TaobaologisticswmspackageentryorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmspackageentryorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packageentryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmspackageentryorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmspackageentryorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmPackageEntryOrderRequest is ConfirmPackageEntryOrderRequest Setter
// 请求
func (r *TaobaologisticswmspackageentryorderconfirmAPIRequest) SetConfirmPackageEntryOrderRequest(_confirmPackageEntryOrderRequest *ConfirmPackageOrderRequest) error {
	r._confirmPackageEntryOrderRequest = _confirmPackageEntryOrderRequest
	r.Set("confirm_package_entry_order_request", _confirmPackageEntryOrderRequest)
	return nil
}

// GetConfirmPackageEntryOrderRequest ConfirmPackageEntryOrderRequest Getter
func (r TaobaologisticswmspackageentryorderconfirmAPIRequest) GetConfirmPackageEntryOrderRequest() *ConfirmPackageOrderRequest {
	return r._confirmPackageEntryOrderRequest
}
