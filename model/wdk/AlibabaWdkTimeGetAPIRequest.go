package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTimeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.time.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTimeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
