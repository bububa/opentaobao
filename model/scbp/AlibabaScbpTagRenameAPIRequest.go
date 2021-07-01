package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重命名关键词分组 API请求
alibaba.scbp.tag.rename

重命名关键词分组
*/
type AlibabaScbpTagRenameAPIRequest struct {
    model.Params
    // 需要重命名的关键词分组名
    _tagName   string
    // 新分组名
    _newTagName   string
}

// 初始化AlibabaScbpTagRenameAPIRequest对象
func NewAlibabaScbpTagRenameRequest() *AlibabaScbpTagRenameAPIRequest{
    return &AlibabaScbpTagRenameAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagRenameAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.rename"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagRenameAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagName Setter
// 需要重命名的关键词分组名
func (r *AlibabaScbpTagRenameAPIRequest) SetTagName(_tagName string) error {
    r._tagName = _tagName
    r.Set("tag_name", _tagName)
    return nil
}

// TagName Getter
func (r AlibabaScbpTagRenameAPIRequest) GetTagName() string {
    return r._tagName
}
// NewTagName Setter
// 新分组名
func (r *AlibabaScbpTagRenameAPIRequest) SetNewTagName(_newTagName string) error {
    r._newTagName = _newTagName
    r.Set("new_tag_name", _newTagName)
    return nil
}

// NewTagName Getter
func (r AlibabaScbpTagRenameAPIRequest) GetNewTagName() string {
    return r._newTagName
}
