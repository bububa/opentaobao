package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTimeGetAPIRequest 获得当前系统时间 API请求
// alibaba.wdk.time.get
//
// 获得当前系统时间
type AlibabaWdkTimeGetAPIRequest struct {
	model.Params
}

// NewAlibabaWdkTimeGetRequest 初始化AlibabaWdkTimeGetAPIRequest对象
func NewAlibabaWdkTimeGetRequest() *AlibabaWdkTimeGetAPIRequest {
	return &AlibabaWdkTimeGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTimeGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTimeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.time.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTimeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTimeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaWdkTimeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTimeGetRequest()
	},
}

// GetAlibabaWdkTimeGetRequest 从 sync.Pool 获取 AlibabaWdkTimeGetAPIRequest
func GetAlibabaWdkTimeGetAPIRequest() *AlibabaWdkTimeGetAPIRequest {
	return poolAlibabaWdkTimeGetAPIRequest.Get().(*AlibabaWdkTimeGetAPIRequest)
}

// ReleaseAlibabaWdkTimeGetAPIRequest 将 AlibabaWdkTimeGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTimeGetAPIRequest(v *AlibabaWdkTimeGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkTimeGetAPIRequest.Put(v)
}
