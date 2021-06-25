package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建关键词分组 APIRequest
alibaba.scbp.tag.add

创建关键词分组
*/
type AlibabaScbpTagAddRequest struct {
    model.Params

    // 分组名称，最多允许创建100个
    tagName   string 

}

func NewAlibabaScbpTagAddRequest() *AlibabaScbpTagAddRequest{
    return &AlibabaScbpTagAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTagAddRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.add"
}

func (r AlibabaScbpTagAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTagAddRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

func (r AlibabaScbpTagAddRequest) GetTagName() string {
    return r.tagName
}

