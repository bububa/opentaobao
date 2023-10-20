package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomercsyncAPIRequest 阿里房产图文草稿信息同步 API请求
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
type AlibabaalihousenewhomercsyncAPIRequest struct {
	model.Params
	// 图文内容
	_rc *RichContentDraftDto
}

// NewAlibabaalihousenewhomercsyncRequest 初始化AlibabaalihousenewhomercsyncAPIRequest对象
func NewAlibabaalihousenewhomercsyncRequest() *AlibabaalihousenewhomercsyncAPIRequest {
	return &AlibabaalihousenewhomercsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomercsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.rc.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomercsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomercsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRc is Rc Setter
// 图文内容
func (r *AlibabaalihousenewhomercsyncAPIRequest) SetRc(_rc *RichContentDraftDto) error {
	r._rc = _rc
	r.Set("rc", _rc)
	return nil
}

// GetRc Rc Getter
func (r AlibabaalihousenewhomercsyncAPIRequest) GetRc() *RichContentDraftDto {
	return r._rc
}
