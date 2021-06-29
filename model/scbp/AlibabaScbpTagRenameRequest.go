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
type AlibabaScbpTagRenameRequest struct {
    model.Params
    // 需要重命名的关键词分组名
    _tagName   string
    // 新分组名
    _newTagName   string
}

// 初始化AlibabaScbpTagRenameRequest对象
func NewAlibabaScbpTagRenameRequest() *AlibabaScbpTagRenameRequest{
    return &AlibabaScbpTagRenameRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagRenameRequest) GetApiMethodName() string {
    return "alibaba.scbp.tag.rename"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagRenameRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagName Setter
// 需要重命名的关键词分组名
func (r *AlibabaScbpTagRenameRequest) SetTagName(_tagName string) error {
    r._tagName = _tagName
    r.Set("tag_name", _tagName)
    return nil
}

// TagName Getter
func (r AlibabaScbpTagRenameRequest) GetTagName() string {
    return r._tagName
}
// NewTagName Setter
// 新分组名
func (r *AlibabaScbpTagRenameRequest) SetNewTagName(_newTagName string) error {
    r._newTagName = _newTagName
    r.Set("new_tag_name", _newTagName)
    return nil
}

// NewTagName Getter
func (r AlibabaScbpTagRenameRequest) GetNewTagName() string {
    return r._newTagName
}
