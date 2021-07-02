package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBoxPostBackBoxAPIRequest RT收箱回传 API请求
// alibaba.wdk.fulfill.box.post.back.box
//
// RT收箱后，信息同步履约，履约同通知UMS 容器管理
type AlibabaWdkFulfillBoxPostBackBoxAPIRequest struct {
	model.Params
	// RT收箱回传请求参数
	_returnBoxContainerRequest *ReturnBoxContainerRequest
}

// NewAlibabaWdkFulfillBoxPostBackBoxRequest 初始化AlibabaWdkFulfillBoxPostBackBoxAPIRequest对象
func NewAlibabaWdkFulfillBoxPostBackBoxRequest() *AlibabaWdkFulfillBoxPostBackBoxAPIRequest {
	return &AlibabaWdkFulfillBoxPostBackBoxAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBoxPostBackBoxAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.box.post.back.box"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBoxPostBackBoxAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReturnBoxContainerRequest is ReturnBoxContainerRequest Setter
// RT收箱回传请求参数
func (r *AlibabaWdkFulfillBoxPostBackBoxAPIRequest) SetReturnBoxContainerRequest(_returnBoxContainerRequest *ReturnBoxContainerRequest) error {
	r._returnBoxContainerRequest = _returnBoxContainerRequest
	r.Set("return_box_container_request", _returnBoxContainerRequest)
	return nil
}

// GetReturnBoxContainerRequest ReturnBoxContainerRequest Getter
func (r AlibabaWdkFulfillBoxPostBackBoxAPIRequest) GetReturnBoxContainerRequest() *ReturnBoxContainerRequest {
	return r._returnBoxContainerRequest
}
