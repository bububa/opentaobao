package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeVideoSyncAPIRequest
视频草稿信息同步 API请求
alibaba.alihouse.newhome.video.sync

接收视频信息记录 */
type AlibabaAlihouseNewhomeVideoSyncAPIRequest struct {
	model.Params
	// 草稿对下
	_video *VideoDraftDto
}

// New
