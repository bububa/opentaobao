package ott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottplayservicegetplayurlAPIRequest 获取播放串地址 API请求
// youku.ott.playservice.getplayurl
//
// 获取播放串地址服务
type YoukuottplayservicegetplayurlAPIRequest struct {
	model.Params
	// 优酷账号登录态
	_yktk string
	// 账号登录态
	_havanaToken string
	// 系统信息
	_systemInfo string
	// 视频ID
	_videoId int64
}

// NewYoukuottplayservicegetplayurlRequest 初始化YoukuottplayservicegetplayurlAPIRequest对象
func NewYoukuottplayservicegetplayurlRequest() *YoukuottplayservicegetplayurlAPIRequest {
	return &YoukuottplayservicegetplayurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottplayservicegetplayurlAPIRequest) GetApiMethodName() string {
	return "youku.ott.playservice.getplayurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottplayservicegetplayurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottplayservicegetplayurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetYktk is Yktk Setter
// 优酷账号登录态
func (r *YoukuottplayservicegetplayurlAPIRequest) SetYktk(_yktk string) error {
	r._yktk = _yktk
	r.Set("yktk", _yktk)
	return nil
}

// GetYktk Yktk Getter
func (r YoukuottplayservicegetplayurlAPIRequest) GetYktk() string {
	return r._yktk
}

// SetHavanaToken is HavanaToken Setter
// 账号登录态
func (r *YoukuottplayservicegetplayurlAPIRequest) SetHavanaToken(_havanaToken string) error {
	r._havanaToken = _havanaToken
	r.Set("havana_token", _havanaToken)
	return nil
}

// GetHavanaToken HavanaToken Getter
func (r YoukuottplayservicegetplayurlAPIRequest) GetHavanaToken() string {
	return r._havanaToken
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *YoukuottplayservicegetplayurlAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r YoukuottplayservicegetplayurlAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetVideoId is VideoId Setter
// 视频ID
func (r *YoukuottplayservicegetplayurlAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r YoukuottplayservicegetplayurlAPIRequest) GetVideoId() int64 {
	return r._videoId
}
