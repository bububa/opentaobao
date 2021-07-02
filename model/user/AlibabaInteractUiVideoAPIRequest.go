package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractUiVideoAPIRequest 视频播放 API请求
// alibaba.interact.ui.video
//
// Weex页面播放视频
type AlibabaInteractUiVideoAPIRequest struct {
	model.Params
	// 仅作鉴权使用，没有实际数据传输
	_unnamed string
}

// NewAlibabaInteractUiVideoRequest 初始化AlibabaInteractUiVideoAPIRequest对象
func NewAlibabaInteractUiVideoRequest() *AlibabaInteractUiVideoAPIRequest {
	return &AlibabaInteractUiVideoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractUiVideoAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.ui.video"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractUiVideoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Unnamed Setter
// 仅作鉴权使用，没有实际数据传输
func (r *AlibabaInteractUiVideoAPIRequest) SetUnnamed(_unnamed string) error {
	r._unnamed = _unnamed
	r.Set("unnamed", _unnamed)
	return nil
}

// Get Unnamed Getter
func (r AlibabaInteractUiVideoAPIRequest) GetUnnamed() string {
	return r._unnamed
}
