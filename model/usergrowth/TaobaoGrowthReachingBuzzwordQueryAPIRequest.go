package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingbuzzwordqueryAPIRequest 淘宝热词榜单数据查询接口 API请求
// taobao.growth.reaching.buzzword.query
//
// 查询淘宝热词榜单数据
type TaobaogrowthreachingbuzzwordqueryAPIRequest struct {
	model.Params
	// 媒体站点类型
	_siteId string
	// 设备信息
	_deviceIds *DeviceIdParam
	// 需要的数据量，如果产出的数据不足，会返回错误代码：data insufficient
	_wantedSize int64
	// 数据偏移位置
	_dataOffset int64
}

// NewTaobaogrowthreachingbuzzwordqueryRequest 初始化TaobaogrowthreachingbuzzwordqueryAPIRequest对象
func NewTaobaogrowthreachingbuzzwordqueryRequest() *TaobaogrowthreachingbuzzwordqueryAPIRequest {
	return &TaobaogrowthreachingbuzzwordqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.buzzword.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSiteId is SiteId Setter
// 媒体站点类型
func (r *TaobaogrowthreachingbuzzwordqueryAPIRequest) SetSiteId(_siteId string) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetSiteId() string {
	return r._siteId
}

// SetDeviceIds is DeviceIds Setter
// 设备信息
func (r *TaobaogrowthreachingbuzzwordqueryAPIRequest) SetDeviceIds(_deviceIds *DeviceIdParam) error {
	r._deviceIds = _deviceIds
	r.Set("device_ids", _deviceIds)
	return nil
}

// GetDeviceIds DeviceIds Getter
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetDeviceIds() *DeviceIdParam {
	return r._deviceIds
}

// SetWantedSize is WantedSize Setter
// 需要的数据量，如果产出的数据不足，会返回错误代码：data insufficient
func (r *TaobaogrowthreachingbuzzwordqueryAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}

// SetDataOffset is DataOffset Setter
// 数据偏移位置
func (r *TaobaogrowthreachingbuzzwordqueryAPIRequest) SetDataOffset(_dataOffset int64) error {
	r._dataOffset = _dataOffset
	r.Set("data_offset", _dataOffset)
	return nil
}

// GetDataOffset DataOffset Getter
func (r TaobaogrowthreachingbuzzwordqueryAPIRequest) GetDataOffset() int64 {
	return r._dataOffset
}
