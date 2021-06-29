package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建关键词分组 API请求
alibaba.scbp.tag.add

创建关键词分组
*/
type AlibabaScbpTagAddRequest struct {
    model.Params
    // 分组名称，最多允许创建100个
    tagName   string
}

// 初始化AlibabaScbpTagAddRequest对象
func NewAlibabaScbpTagAddRequest() *AlibabaScbpTagAddRequest{
    return &AlibabaScbpTagAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagAddRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagName Setter
// 分组名称，最多允许创建100个
func (r *AlibabaScbpTagAddRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

// TagName Getter
func (r AlibabaScbpTagAddRequest) GetTagName() string {
    return r.tagName
}
