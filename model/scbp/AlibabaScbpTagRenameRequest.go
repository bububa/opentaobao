package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重命名关键词分组 APIRequest
alibaba.scbp.tag.rename

重命名关键词分组
*/
type AlibabaScbpTagRenameRequest struct {
    model.Params

    // 需要重命名的关键词分组名
    tagName   string 

    // 新分组名
    newTagName   string 

}

func NewAlibabaScbpTagRenameRequest() *AlibabaScbpTagRenameRequest{
    return &AlibabaScbpTagRenameRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTagRenameRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.rename"
}

func (r AlibabaScbpTagRenameRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTagRenameRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

func (r AlibabaScbpTagRenameRequest) GetTagName() string {
    return r.tagName
}

func (r *AlibabaScbpTagRenameRequest) SetNewTagName(newTagName string) error {
    r.newTagName = newTagName
    r.Set("new_tag_name", newTagName)
    return nil
}

func (r AlibabaScbpTagRenameRequest) GetNewTagName() string {
    return r.newTagName
}

