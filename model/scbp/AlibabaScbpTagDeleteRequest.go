package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词分组 API请求
alibaba.scbp.tag.delete

删除关键词分组
*/
type AlibabaScbpTagDeleteRequest struct {
    model.Params
    // 关键词分组名
    tagName   string
}

// 初始化AlibabaScbpTagDeleteRequest对象
func NewAlibabaScbpTagDeleteRequest() *AlibabaScbpTagDeleteRequest{
    return &AlibabaScbpTagDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagDeleteRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagName Setter
// 关键词分组名
func (r *AlibabaScbpTagDeleteRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

// TagName Getter
func (r AlibabaScbpTagDeleteRequest) GetTagName() string {
    return r.tagName
}
