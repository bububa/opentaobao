package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillBoxPostBackBoxAPIRequest) Reset() {
	r._returnBoxContainerRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBoxPostBackBoxAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.box.post.back.box"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBoxPostBackBoxAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillBoxPostBackBoxAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkFulfillBoxPostBackBoxAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillBoxPostBackBoxRequest()
	},
}

// GetAlibabaWdkFulfillBoxPostBackBoxRequest 从 sync.Pool 获取 AlibabaWdkFulfillBoxPostBackBoxAPIRequest
func GetAlibabaWdkFulfillBoxPostBackBoxAPIRequest() *AlibabaWdkFulfillBoxPostBackBoxAPIRequest {
	return poolAlibabaWdkFulfillBoxPostBackBoxAPIRequest.Get().(*AlibabaWdkFulfillBoxPostBackBoxAPIRequest)
}

// ReleaseAlibabaWdkFulfillBoxPostBackBoxAPIRequest 将 AlibabaWdkFulfillBoxPostBackBoxAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillBoxPostBackBoxAPIRequest(v *AlibabaWdkFulfillBoxPostBackBoxAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillBoxPostBackBoxAPIRequest.Put(v)
}
