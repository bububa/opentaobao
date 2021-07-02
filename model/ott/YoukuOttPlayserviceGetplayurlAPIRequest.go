package ott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPlayserviceGetplayurlAPIRequest 获取播放串地址 API请求
// youku.ott.playservice.getplayurl
//
// 获取播放串地址服务
type YoukuOttPlayserviceGetplayurlAPIRequest struct {
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

// NewYoukuOttPlayserviceGetplayurlRequest 初始化YoukuOttPlayserviceGetplayurlAPIRequest对象
func NewYoukuOttPlayserviceGetplayurlRequest() *YoukuOttPlayserviceGetplayurlAPIRequest {
	return &YoukuOttPlayserviceGetplayurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPlayserviceGetplayurlAPIRequest) GetApiMethodName() string {
	return "youku.ott.playservice.getplayurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPlayserviceGetplayurlAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetYktk is Yktk Setter
// 优酷账号登录态
func (r *YoukuOttPlayserviceGetplayurlAPIRequest) SetYktk(_yktk string) error {
	r._yktk = _yktk
	r.Set("yktk", _yktk)
	return nil
}

// GetYktk Yktk Getter
func (r YoukuOttPlayserviceGetplayurlAPIRequest) GetYktk() string {
	return r._yktk
}

// SetHavanaToken is HavanaToken Setter
// 账号登录态
func (r *YoukuOttPlayserviceGetplayurlAPIRequest) SetHavanaToken(_havanaToken string) error {
	r._havanaToken = _havanaToken
	r.Set("havana_token", _havanaToken)
	return nil
}

// GetHavanaToken HavanaToken Getter
func (r YoukuOttPlayserviceGetplayurlAPIRequest) GetHavanaToken() string {
	return r._havanaToken
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *YoukuOttPlayserviceGetplayurlAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r YoukuOttPlayserviceGetplayurlAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetVideoId is VideoId Setter
// 视频ID
func (r *YoukuOttPlayserviceGetplayurlAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r YoukuOttPlayserviceGetplayurlAPIRequest) GetVideoId() int64 {
	return r._videoId
}
