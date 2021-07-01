package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgVideoUploadAPIRequest
提供优酷的短视频入淘API API请求
alibaba.baichuan.ctg.video.upload

提供优酷的短视频入淘API */
type AlibabaBaichuanCtgVideoUploadAPIRequest struct {
	model.Params
	// app
	_app string
	// type
	_type string
	// 优酷道长绑定的淘宝账号ID
	_tbUid string
	// 视频VID，若为多个视频，则支持分组上传多个VID
	_videoId string
	// 作者名称
	_ownerName string
	// 发布时间
	_publishTime string
	// 上传时间
	_uploadTime string
	// 视频标题
	_videoTitle string
	// 视频描述
	_videoInfo string
	// 视频的分类ID，目前是优酷的分类ID
	_videoCategory string
	// 视频标签
	_videoTag string
	// 视频的平台来源，如，优酷
	_source string
}

// New
