package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkHrworkbenchCdporgsQueryAPIRequest
homs人力资源组织树查询接口 API请求
alibaba.wdk.hrworkbench.cdporgs.query

提供查询homs人力组织树的接口，按照商家做权限隔离。 */
type AlibabaWdkHrworkbenchCdporgsQueryAPIRequest struct {
	model.Params
}

// NewAlibabaWdkHrworkbenchCdporgsQueryRequest 初始化AlibabaWdkHrworkbenchCdporgsQueryAPIRequest对象
func NewAlibabaWdkHrworkbenchCdporgsQueryRequest() *AlibabaWdkHrworkbenchCdporgsQueryAPIRequest {
	return &AlibabaWdkHrworkbenchCdporgsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchCdporgsQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.hrworkbench.cdporgs.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchCdporgsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
