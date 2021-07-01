package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分组信息获取 API请求
alibaba.icbu.product.group.get

分组信息获取
*/
type AlibabaIcbuProductGroupGetAPIRequest struct {
    model.Params
    // 分组ID，传-1获得所有一级分组
    _groupId   int64
    // 补充信息
    _extraContext   string
}

// 初始化AlibabaIcbuProductGroupGetAPIRequest对象
func NewAlibabaIcbuProductGroupGetRequest() *AlibabaIcbuProductGroupGetAPIRequest{
    return &AlibabaIcbuProductGroupGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGroupGetAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.group.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGroupGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 分组ID，传-1获得所有一级分组
func (r *AlibabaIcbuProductGroupGetAPIRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuProductGroupGetAPIRequest) GetGroupId() int64 {
    return r._groupId
}
// ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductGroupGetAPIRequest) SetExtraContext(_extraContext string) error {
    r._extraContext = _extraContext
    r.Set("extra_context", _extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuProductGroupGetAPIRequest) GetExtraContext() string {
    return r._extraContext
}
