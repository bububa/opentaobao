package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品分组 API请求
alibaba.icbu.product.group.add

增加商品分组
*/
type AlibabaIcbuProductGroupAddRequest struct {
    model.Params
    // 分组名称
    _groupName   string
    // 上级分组ID，如果建立顶级分组设为-1
    _parentId   int64
    // 补充信息，如isv id
    _extraContext   string
}

// 初始化AlibabaIcbuProductGroupAddRequest对象
func NewAlibabaIcbuProductGroupAddRequest() *AlibabaIcbuProductGroupAddRequest{
    return &AlibabaIcbuProductGroupAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGroupAddRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.group.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 分组名称
func (r *AlibabaIcbuProductGroupAddRequest) SetGroupName(_groupName string) error {
    r._groupName = _groupName
    r.Set("group_name", _groupName)
    return nil
}

// GroupName Getter
func (r AlibabaIcbuProductGroupAddRequest) GetGroupName() string {
    return r._groupName
}
// ParentId Setter
// 上级分组ID，如果建立顶级分组设为-1
func (r *AlibabaIcbuProductGroupAddRequest) SetParentId(_parentId int64) error {
    r._parentId = _parentId
    r.Set("parent_id", _parentId)
    return nil
}

// ParentId Getter
func (r AlibabaIcbuProductGroupAddRequest) GetParentId() int64 {
    return r._parentId
}
// ExtraContext Setter
// 补充信息，如isv id
func (r *AlibabaIcbuProductGroupAddRequest) SetExtraContext(_extraContext string) error {
    r._extraContext = _extraContext
    r.Set("extra_context", _extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuProductGroupAddRequest) GetExtraContext() string {
    return r._extraContext
}
