package usergrowth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthAdMaterialDataSyncAPIRequest 素材投放效果数据回传 API请求
// taobao.usergrowth.ad.material.data.sync
//
// 创意维度广告效果数据回传
type TaobaoUsergrowthAdMaterialDataSyncAPIRequest struct {
	model.Params
	// 素材url
	_materialUrl string
	// 统计日期
	_statDate string
	// 文件md5
	_signature string
	// 素材编码
	_materialCode string
	// 其他媒体
	_otherMedia string
	// 资源位
	_adPlacement string
	// 媒体
	_mediaCode string
	// 消耗
	_cost int64
	// 操作系统
	_deviceOs int64
	// 点击
	_click int64
	// 曝光
	_exposure int64
	// 渠道id
	_channelId int64
	// 任务id
	_taskId int64
}

// NewTaobaoUsergrowthAdMaterialDataSyncRequest 初始化TaobaoUsergrowthAdMaterialDataSyncAPIRequest对象
func NewTaobaoUsergrowthAdMaterialDataSyncRequest() *TaobaoUsergrowthAdMaterialDataSyncAPIRequest {
	return &TaobaoUsergrowthAdMaterialDataSyncAPIRequest{
		Params: model.NewParams(13),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) Reset() {
	r._materialUrl = ""
	r._statDate = ""
	r._signature = ""
	r._materialCode = ""
	r._otherMedia = ""
	r._adPlacement = ""
	r._mediaCode = ""
	r._cost = 0
	r._deviceOs = 0
	r._click = 0
	r._exposure = 0
	r._channelId = 0
	r._taskId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.material.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialUrl is MaterialUrl Setter
// 素材url
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetMaterialUrl(_materialUrl string) error {
	r._materialUrl = _materialUrl
	r.Set("material_url", _materialUrl)
	return nil
}

// GetMaterialUrl MaterialUrl Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetMaterialUrl() string {
	return r._materialUrl
}

// SetStatDate is StatDate Setter
// 统计日期
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetStatDate(_statDate string) error {
	r._statDate = _statDate
	r.Set("stat_date", _statDate)
	return nil
}

// GetStatDate StatDate Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetStatDate() string {
	return r._statDate
}

// SetSignature is Signature Setter
// 文件md5
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetSignature() string {
	return r._signature
}

// SetMaterialCode is MaterialCode Setter
// 素材编码
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetMaterialCode(_materialCode string) error {
	r._materialCode = _materialCode
	r.Set("material_code", _materialCode)
	return nil
}

// GetMaterialCode MaterialCode Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetMaterialCode() string {
	return r._materialCode
}

// SetOtherMedia is OtherMedia Setter
// 其他媒体
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetOtherMedia(_otherMedia string) error {
	r._otherMedia = _otherMedia
	r.Set("other_media", _otherMedia)
	return nil
}

// GetOtherMedia OtherMedia Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetOtherMedia() string {
	return r._otherMedia
}

// SetAdPlacement is AdPlacement Setter
// 资源位
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetAdPlacement(_adPlacement string) error {
	r._adPlacement = _adPlacement
	r.Set("ad_placement", _adPlacement)
	return nil
}

// GetAdPlacement AdPlacement Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetAdPlacement() string {
	return r._adPlacement
}

// SetMediaCode is MediaCode Setter
// 媒体
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetMediaCode(_mediaCode string) error {
	r._mediaCode = _mediaCode
	r.Set("media_code", _mediaCode)
	return nil
}

// GetMediaCode MediaCode Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetMediaCode() string {
	return r._mediaCode
}

// SetCost is Cost Setter
// 消耗
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetCost(_cost int64) error {
	r._cost = _cost
	r.Set("cost", _cost)
	return nil
}

// GetCost Cost Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetCost() int64 {
	return r._cost
}

// SetDeviceOs is DeviceOs Setter
// 操作系统
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetDeviceOs(_deviceOs int64) error {
	r._deviceOs = _deviceOs
	r.Set("device_os", _deviceOs)
	return nil
}

// GetDeviceOs DeviceOs Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetDeviceOs() int64 {
	return r._deviceOs
}

// SetClick is Click Setter
// 点击
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetClick(_click int64) error {
	r._click = _click
	r.Set("click", _click)
	return nil
}

// GetClick Click Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetClick() int64 {
	return r._click
}

// SetExposure is Exposure Setter
// 曝光
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetExposure(_exposure int64) error {
	r._exposure = _exposure
	r.Set("exposure", _exposure)
	return nil
}

// GetExposure Exposure Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetExposure() int64 {
	return r._exposure
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetChannelId() int64 {
	return r._channelId
}

// SetTaskId is TaskId Setter
// 任务id
func (r *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoUsergrowthAdMaterialDataSyncAPIRequest) GetTaskId() int64 {
	return r._taskId
}

var poolTaobaoUsergrowthAdMaterialDataSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsergrowthAdMaterialDataSyncRequest()
	},
}

// GetTaobaoUsergrowthAdMaterialDataSyncRequest 从 sync.Pool 获取 TaobaoUsergrowthAdMaterialDataSyncAPIRequest
func GetTaobaoUsergrowthAdMaterialDataSyncAPIRequest() *TaobaoUsergrowthAdMaterialDataSyncAPIRequest {
	return poolTaobaoUsergrowthAdMaterialDataSyncAPIRequest.Get().(*TaobaoUsergrowthAdMaterialDataSyncAPIRequest)
}

// ReleaseTaobaoUsergrowthAdMaterialDataSyncAPIRequest 将 TaobaoUsergrowthAdMaterialDataSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsergrowthAdMaterialDataSyncAPIRequest(v *TaobaoUsergrowthAdMaterialDataSyncAPIRequest) {
	v.Reset()
	poolTaobaoUsergrowthAdMaterialDataSyncAPIRequest.Put(v)
}
