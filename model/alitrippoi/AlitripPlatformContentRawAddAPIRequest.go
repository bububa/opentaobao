package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformcontentrawaddAPIRequest 穷游内容写入接口 API请求
// alitrip.platform.content.raw.add
//
// 穷游内容写入飞猪接口
type AlitripplatformcontentrawaddAPIRequest struct {
	model.Params
	// 写入入参
	_fliggyContentRequest *FliggyContentRequest
}

// NewAlitripplatformcontentrawaddRequest 初始化AlitripplatformcontentrawaddAPIRequest对象
func NewAlitripplatformcontentrawaddRequest() *AlitripplatformcontentrawaddAPIRequest {
	return &AlitripplatformcontentrawaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripplatformcontentrawaddAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.content.raw.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripplatformcontentrawaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripplatformcontentrawaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFliggyContentRequest is FliggyContentRequest Setter
// 写入入参
func (r *AlitripplatformcontentrawaddAPIRequest) SetFliggyContentRequest(_fliggyContentRequest *FliggyContentRequest) error {
	r._fliggyContentRequest = _fliggyContentRequest
	r.Set("fliggy_content_request", _fliggyContentRequest)
	return nil
}

// GetFliggyContentRequest FliggyContentRequest Getter
func (r AlitripplatformcontentrawaddAPIRequest) GetFliggyContentRequest() *FliggyContentRequest {
	return r._fliggyContentRequest
}
