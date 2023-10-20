package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest 包裹入库单确认 API请求
// taobao.logistics.wms.packageentryorder.confirm
//
// 包裹入库单确认
type TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest struct {
	model.Params
	// 请求
	_confirmPackageEntryOrderRequest *ConfirmPackageOrderRequest
}

// NewTaobaoLogisticsWmsPackageentryorderConfirmRequest 初始化TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest对象
func NewTaobaoLogisticsWmsPackageentryorderConfirmRequest() *TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest {
	return &TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) Reset() {
	r._confirmPackageEntryOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packageentryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmPackageEntryOrderRequest is ConfirmPackageEntryOrderRequest Setter
// 请求
func (r *TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) SetConfirmPackageEntryOrderRequest(_confirmPackageEntryOrderRequest *ConfirmPackageOrderRequest) error {
	r._confirmPackageEntryOrderRequest = _confirmPackageEntryOrderRequest
	r.Set("confirm_package_entry_order_request", _confirmPackageEntryOrderRequest)
	return nil
}

// GetConfirmPackageEntryOrderRequest ConfirmPackageEntryOrderRequest Getter
func (r TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) GetConfirmPackageEntryOrderRequest() *ConfirmPackageOrderRequest {
	return r._confirmPackageEntryOrderRequest
}

var poolTaobaoLogisticsWmsPackageentryorderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWmsPackageentryorderConfirmRequest()
	},
}

// GetTaobaoLogisticsWmsPackageentryorderConfirmRequest 从 sync.Pool 获取 TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest
func GetTaobaoLogisticsWmsPackageentryorderConfirmAPIRequest() *TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest {
	return poolTaobaoLogisticsWmsPackageentryorderConfirmAPIRequest.Get().(*TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest)
}

// ReleaseTaobaoLogisticsWmsPackageentryorderConfirmAPIRequest 将 TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWmsPackageentryorderConfirmAPIRequest(v *TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWmsPackageentryorderConfirmAPIRequest.Put(v)
}
