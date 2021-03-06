package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformContentRawAddAPIRequest 穷游内容写入接口 API请求
// alitrip.platform.content.raw.add
//
// 穷游内容写入飞猪接口
type AlitripPlatformContentRawAddAPIRequest struct {
	model.Params
	// 写入入参
	_fliggyContentRequest *FliggyContentRequest
}

// NewAlitripPlatformContentRawAddRequest 初始化AlitripPlatformContentRawAddAPIRequest对象
func NewAlitripPlatformContentRawAddRequest() *AlitripPlatformContentRawAddAPIRequest {
	return &AlitripPlatformContentRawAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformContentRawAddAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.content.raw.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformContentRawAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFliggyContentRequest is FliggyContentRequest Setter
// 写入入参
func (r *AlitripPlatformContentRawAddAPIRequest) SetFliggyContentRequest(_fliggyContentRequest *FliggyContentRequest) error {
	r._fliggyContentRequest = _fliggyContentRequest
	r.Set("fliggy_content_request", _fliggyContentRequest)
	return nil
}

// GetFliggyContentRequest FliggyContentRequest Getter
func (r AlitripPlatformContentRawAddAPIRequest) GetFliggyContentRequest() *FliggyContentRequest {
	return r._fliggyContentRequest
}
