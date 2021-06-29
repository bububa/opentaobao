package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询所有分组 API请求
alibaba.scbp.tag.list

查询所有分组
*/
type AlibabaScbpTagListRequest struct {
    model.Params
}

// 初始化AlibabaScbpTagListRequest对象
func NewAlibabaScbpTagListRequest() *AlibabaScbpTagListRequest{
    return &AlibabaScbpTagListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagListRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
