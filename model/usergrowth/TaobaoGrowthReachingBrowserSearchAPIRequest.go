package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingbrowsersearchAPIRequest 查询搜索关联 API请求
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
type TaobaogrowthreachingbrowsersearchAPIRequest struct {
	model.Params
	// 查询词
	_query string
	// 设备列表
	_deviceIds *DeviceIdParam
	// 需要数量
	_wantedSize int64
}

// NewTaobaogrowthreachingbrowsersearchRequest 初始化TaobaogrowthreachingbrowsersearchAPIRequest对象
func NewTaobaogrowthreachingbrowsersearchRequest() *TaobaogrowthreachingbrowsersearchAPIRequest {
	return &TaobaogrowthreachingbrowsersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogrowthreachingbrowsersearchAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.browser.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogrowthreachingbrowsersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogrowthreachingbrowsersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询词
func (r *TaobaogrowthreachingbrowsersearchAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaogrowthreachingbrowsersearchAPIRequest) GetQuery() string {
	return r._query
}

// SetDeviceIds is DeviceIds Setter
// 设备列表
func (r *TaobaogrowthreachingbrowsersearchAPIRequest) SetDeviceIds(_deviceIds *DeviceIdParam) error {
	r._deviceIds = _deviceIds
	r.Set("device_ids", _deviceIds)
	return nil
}

// GetDeviceIds DeviceIds Getter
func (r TaobaogrowthreachingbrowsersearchAPIRequest) GetDeviceIds() *DeviceIdParam {
	return r._deviceIds
}

// SetWantedSize is WantedSize Setter
// 需要数量
func (r *TaobaogrowthreachingbrowsersearchAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaogrowthreachingbrowsersearchAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}
