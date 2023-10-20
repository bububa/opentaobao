package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillboxpostbackboxAPIRequest RT收箱回传 API请求
// alibaba.wdk.fulfill.box.post.back.box
//
// RT收箱后，信息同步履约，履约同通知UMS 容器管理
type AlibabawdkfulfillboxpostbackboxAPIRequest struct {
	model.Params
	// RT收箱回传请求参数
	_returnBoxContainerRequest *ReturnBoxContainerRequest
}

// NewAlibabawdkfulfillboxpostbackboxRequest 初始化AlibabawdkfulfillboxpostbackboxAPIRequest对象
func NewAlibabawdkfulfillboxpostbackboxRequest() *AlibabawdkfulfillboxpostbackboxAPIRequest {
	return &AlibabawdkfulfillboxpostbackboxAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillboxpostbackboxAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.box.post.back.box"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillboxpostbackboxAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillboxpostbackboxAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReturnBoxContainerRequest is ReturnBoxContainerRequest Setter
// RT收箱回传请求参数
func (r *AlibabawdkfulfillboxpostbackboxAPIRequest) SetReturnBoxContainerRequest(_returnBoxContainerRequest *ReturnBoxContainerRequest) error {
	r._returnBoxContainerRequest = _returnBoxContainerRequest
	r.Set("return_box_container_request", _returnBoxContainerRequest)
	return nil
}

// GetReturnBoxContainerRequest ReturnBoxContainerRequest Getter
func (r AlibabawdkfulfillboxpostbackboxAPIRequest) GetReturnBoxContainerRequest() *ReturnBoxContainerRequest {
	return r._returnBoxContainerRequest
}
