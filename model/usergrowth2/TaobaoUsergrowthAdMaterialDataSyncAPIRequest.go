package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthadmaterialdatasyncAPIRequest 素材投放效果数据回传 API请求
// taobao.usergrowth.ad.material.data.sync
//
// 创意维度广告效果数据回传
type TaobaousergrowthadmaterialdatasyncAPIRequest struct {
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

// NewTaobaousergrowthadmaterialdatasyncRequest 初始化TaobaousergrowthadmaterialdatasyncAPIRequest对象
func NewTaobaousergrowthadmaterialdatasyncRequest() *TaobaousergrowthadmaterialdatasyncAPIRequest {
	return &TaobaousergrowthadmaterialdatasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.material.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialUrl is MaterialUrl Setter
// 素材url
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetMaterialUrl(_materialUrl string) error {
	r._materialUrl = _materialUrl
	r.Set("material_url", _materialUrl)
	return nil
}

// GetMaterialUrl MaterialUrl Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetMaterialUrl() string {
	return r._materialUrl
}

// SetStatDate is StatDate Setter
// 统计日期
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetStatDate(_statDate string) error {
	r._statDate = _statDate
	r.Set("stat_date", _statDate)
	return nil
}

// GetStatDate StatDate Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetStatDate() string {
	return r._statDate
}

// SetSignature is Signature Setter
// 文件md5
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetSignature() string {
	return r._signature
}

// SetMaterialCode is MaterialCode Setter
// 素材编码
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetMaterialCode(_materialCode string) error {
	r._materialCode = _materialCode
	r.Set("material_code", _materialCode)
	return nil
}

// GetMaterialCode MaterialCode Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetMaterialCode() string {
	return r._materialCode
}

// SetOtherMedia is OtherMedia Setter
// 其他媒体
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetOtherMedia(_otherMedia string) error {
	r._otherMedia = _otherMedia
	r.Set("other_media", _otherMedia)
	return nil
}

// GetOtherMedia OtherMedia Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetOtherMedia() string {
	return r._otherMedia
}

// SetAdPlacement is AdPlacement Setter
// 资源位
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetAdPlacement(_adPlacement string) error {
	r._adPlacement = _adPlacement
	r.Set("ad_placement", _adPlacement)
	return nil
}

// GetAdPlacement AdPlacement Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetAdPlacement() string {
	return r._adPlacement
}

// SetMediaCode is MediaCode Setter
// 媒体
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetMediaCode(_mediaCode string) error {
	r._mediaCode = _mediaCode
	r.Set("media_code", _mediaCode)
	return nil
}

// GetMediaCode MediaCode Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetMediaCode() string {
	return r._mediaCode
}

// SetCost is Cost Setter
// 消耗
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetCost(_cost int64) error {
	r._cost = _cost
	r.Set("cost", _cost)
	return nil
}

// GetCost Cost Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetCost() int64 {
	return r._cost
}

// SetDeviceOs is DeviceOs Setter
// 操作系统
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetDeviceOs(_deviceOs int64) error {
	r._deviceOs = _deviceOs
	r.Set("device_os", _deviceOs)
	return nil
}

// GetDeviceOs DeviceOs Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetDeviceOs() int64 {
	return r._deviceOs
}

// SetClick is Click Setter
// 点击
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetClick(_click int64) error {
	r._click = _click
	r.Set("click", _click)
	return nil
}

// GetClick Click Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetClick() int64 {
	return r._click
}

// SetExposure is Exposure Setter
// 曝光
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetExposure(_exposure int64) error {
	r._exposure = _exposure
	r.Set("exposure", _exposure)
	return nil
}

// GetExposure Exposure Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetExposure() int64 {
	return r._exposure
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetChannelId() int64 {
	return r._channelId
}

// SetTaskId is TaskId Setter
// 任务id
func (r *TaobaousergrowthadmaterialdatasyncAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaousergrowthadmaterialdatasyncAPIRequest) GetTaskId() int64 {
	return r._taskId
}
