package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalitempublishAPIRequest 三方入驻-开通服务 API请求
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
type AlibabaalihealthmedicalitempublishAPIRequest struct {
	model.Params
	// 请求
	_request1 *ItemPublishRequest
}

// NewAlibabaalihealthmedicalitempublishRequest 初始化AlibabaalihealthmedicalitempublishAPIRequest对象
func NewAlibabaalihealthmedicalitempublishRequest() *AlibabaalihealthmedicalitempublishAPIRequest {
	return &AlibabaalihealthmedicalitempublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalitempublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.item.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalitempublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalitempublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest1 is Request1 Setter
// 请求
func (r *AlibabaalihealthmedicalitempublishAPIRequest) SetRequest1(_request1 *ItemPublishRequest) error {
	r._request1 = _request1
	r.Set("request1", _request1)
	return nil
}

// GetRequest1 Request1 Getter
func (r AlibabaalihealthmedicalitempublishAPIRequest) GetRequest1() *ItemPublishRequest {
	return r._request1
}
