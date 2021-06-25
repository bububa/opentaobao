package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词分组 APIRequest
alibaba.scbp.tag.delete

删除关键词分组
*/
type AlibabaScbpTagDeleteRequest struct {
    model.Params

    // 关键词分组名
    tagName   string 

}

func NewAlibabaScbpTagDeleteRequest() *AlibabaScbpTagDeleteRequest{
    return &AlibabaScbpTagDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTagDeleteRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.delete"
}

func (r AlibabaScbpTagDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTagDeleteRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

func (r AlibabaScbpTagDeleteRequest) GetTagName() string {
    return r.tagName
}

