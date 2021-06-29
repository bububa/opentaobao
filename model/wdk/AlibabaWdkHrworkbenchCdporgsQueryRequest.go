package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
homs人力资源组织树查询接口 API请求
alibaba.wdk.hrworkbench.cdporgs.query

提供查询homs人力组织树的接口，按照商家做权限隔离。
*/
type AlibabaWdkHrworkbenchCdporgsQueryRequest struct {
    model.Params
}

// 初始化AlibabaWdkHrworkbenchCdporgsQueryRequest对象
func NewAlibabaWdkHrworkbenchCdporgsQueryRequest() *AlibabaWdkHrworkbenchCdporgsQueryRequest{
    return &AlibabaWdkHrworkbenchCdporgsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchCdporgsQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.cdporgs.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchCdporgsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
