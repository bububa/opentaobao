package ott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPlayserviceGetplayurlAPIRequest
获取播放串地址 API请求
youku.ott.playservice.getplayurl

获取播放串地址服务 */
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

// New
