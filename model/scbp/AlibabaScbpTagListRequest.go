package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询所有分组 APIRequest
alibaba.scbp.tag.list

查询所有分组
*/
type AlibabaScbpTagListRequest struct {
    model.Params

}

func NewAlibabaScbpTagListRequest() *AlibabaScbpTagListRequest{
    return &AlibabaScbpTagListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTagListRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.list"
}

func (r AlibabaScbpTagListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


