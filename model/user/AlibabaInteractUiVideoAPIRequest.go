package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractUiVideoAPIRequest
视频播放 API请求
alibaba.interact.ui.video

Weex页面播放视频 */
type AlibabaInteractUiVideoAPIRequest struct {
	model.Params
	// 仅作鉴权使用，没有实际数据传输
	_unnamed string
}

// New
