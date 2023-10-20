package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageentryorderPullAPIRequest 包裹入库单拉单 API请求
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
type TaobaoLogisticsWmsPackageentryorderPullAPIRequest struct {
	model.Params
	// 请求
	_pullPackageEntryOrderRequest *PullPackageOrderRequest
}

// NewTaobaoLogisticsWmsPackageentryorderPullRequest 初始化TaobaoLogisticsWmsPackageentryorderPullAPIRequest对象
func NewTaobaoLogisticsWmsPackageentryorderPullRequest() *TaobaoLogisticsWmsPackageentryorderPullAPIRequest {
	return &TaobaoLogisticsWmsPackageentryorderPullAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWmsPackageentryorderPullAPIRequest) Reset() {
	r._pullPackageEntryOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsPackageentryorderPullAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packageentryorder.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsPackageentryorderPullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsPackageentryorderPullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPullPackageEntryOrderRequest is PullPackageEntryOrderRequest Setter
// 请求
func (r *TaobaoLogisticsWmsPackageentryorderPullAPIRequest) SetPullPackageEntryOrderRequest(_pullPackageEntryOrderRequest *PullPackageOrderRequest) error {
	r._pullPackageEntryOrderRequest = _pullPackageEntryOrderRequest
	r.Set("pull_package_entry_order_request", _pullPackageEntryOrderRequest)
	return nil
}

// GetPullPackageEntryOrderRequest PullPackageEntryOrderRequest Getter
func (r TaobaoLogisticsWmsPackageentryorderPullAPIRequest) GetPullPackageEntryOrderRequest() *PullPackageOrderRequest {
	return r._pullPackageEntryOrderRequest
}

var poolTaobaoLogisticsWmsPackageentryorderPullAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWmsPackageentryorderPullRequest()
	},
}

// GetTaobaoLogisticsWmsPackageentryorderPullRequest 从 sync.Pool 获取 TaobaoLogisticsWmsPackageentryorderPullAPIRequest
func GetTaobaoLogisticsWmsPackageentryorderPullAPIRequest() *TaobaoLogisticsWmsPackageentryorderPullAPIRequest {
	return poolTaobaoLogisticsWmsPackageentryorderPullAPIRequest.Get().(*TaobaoLogisticsWmsPackageentryorderPullAPIRequest)
}

// ReleaseTaobaoLogisticsWmsPackageentryorderPullAPIRequest 将 TaobaoLogisticsWmsPackageentryorderPullAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWmsPackageentryorderPullAPIRequest(v *TaobaoLogisticsWmsPackageentryorderPullAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWmsPackageentryorderPullAPIRequest.Put(v)
}
