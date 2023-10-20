package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswmspackageentryorderpullAPIRequest 包裹入库单拉单 API请求
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
type TaobaologisticswmspackageentryorderpullAPIRequest struct {
	model.Params
	// 请求
	_pullPackageEntryOrderRequest *PullPackageOrderRequest
}

// NewTaobaologisticswmspackageentryorderpullRequest 初始化TaobaologisticswmspackageentryorderpullAPIRequest对象
func NewTaobaologisticswmspackageentryorderpullRequest() *TaobaologisticswmspackageentryorderpullAPIRequest {
	return &TaobaologisticswmspackageentryorderpullAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswmspackageentryorderpullAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packageentryorder.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswmspackageentryorderpullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswmspackageentryorderpullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPullPackageEntryOrderRequest is PullPackageEntryOrderRequest Setter
// 请求
func (r *TaobaologisticswmspackageentryorderpullAPIRequest) SetPullPackageEntryOrderRequest(_pullPackageEntryOrderRequest *PullPackageOrderRequest) error {
	r._pullPackageEntryOrderRequest = _pullPackageEntryOrderRequest
	r.Set("pull_package_entry_order_request", _pullPackageEntryOrderRequest)
	return nil
}

// GetPullPackageEntryOrderRequest PullPackageEntryOrderRequest Getter
func (r TaobaologisticswmspackageentryorderpullAPIRequest) GetPullPackageEntryOrderRequest() *PullPackageOrderRequest {
	return r._pullPackageEntryOrderRequest
}
