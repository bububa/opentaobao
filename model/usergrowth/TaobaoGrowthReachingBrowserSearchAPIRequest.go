package usergrowth

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingBrowserSearchAPIRequest 查询搜索关联 API请求
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
type TaobaoGrowthReachingBrowserSearchAPIRequest struct {
	model.Params
	// 查询词
	_query string
	// 设备列表
	_deviceIds *DeviceIdParam
	// 需要数量
	_wantedSize int64
}

// NewTaobaoGrowthReachingBrowserSearchRequest 初始化TaobaoGrowthReachingBrowserSearchAPIRequest对象
func NewTaobaoGrowthReachingBrowserSearchRequest() *TaobaoGrowthReachingBrowserSearchAPIRequest {
	return &TaobaoGrowthReachingBrowserSearchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGrowthReachingBrowserSearchAPIRequest) Reset() {
	r._query = ""
	r._deviceIds = nil
	r._wantedSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGrowthReachingBrowserSearchAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.browser.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGrowthReachingBrowserSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGrowthReachingBrowserSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询词
func (r *TaobaoGrowthReachingBrowserSearchAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoGrowthReachingBrowserSearchAPIRequest) GetQuery() string {
	return r._query
}

// SetDeviceIds is DeviceIds Setter
// 设备列表
func (r *TaobaoGrowthReachingBrowserSearchAPIRequest) SetDeviceIds(_deviceIds *DeviceIdParam) error {
	r._deviceIds = _deviceIds
	r.Set("device_ids", _deviceIds)
	return nil
}

// GetDeviceIds DeviceIds Getter
func (r TaobaoGrowthReachingBrowserSearchAPIRequest) GetDeviceIds() *DeviceIdParam {
	return r._deviceIds
}

// SetWantedSize is WantedSize Setter
// 需要数量
func (r *TaobaoGrowthReachingBrowserSearchAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaoGrowthReachingBrowserSearchAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}

var poolTaobaoGrowthReachingBrowserSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGrowthReachingBrowserSearchRequest()
	},
}

// GetTaobaoGrowthReachingBrowserSearchRequest 从 sync.Pool 获取 TaobaoGrowthReachingBrowserSearchAPIRequest
func GetTaobaoGrowthReachingBrowserSearchAPIRequest() *TaobaoGrowthReachingBrowserSearchAPIRequest {
	return poolTaobaoGrowthReachingBrowserSearchAPIRequest.Get().(*TaobaoGrowthReachingBrowserSearchAPIRequest)
}

// ReleaseTaobaoGrowthReachingBrowserSearchAPIRequest 将 TaobaoGrowthReachingBrowserSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoGrowthReachingBrowserSearchAPIRequest(v *TaobaoGrowthReachingBrowserSearchAPIRequest) {
	v.Reset()
	poolTaobaoGrowthReachingBrowserSearchAPIRequest.Put(v)
}
