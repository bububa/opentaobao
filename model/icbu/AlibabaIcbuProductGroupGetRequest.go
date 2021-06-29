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
type AlibabaIcbuProductGroupGetRequest struct {
    model.Params
    // 分组ID，传-1获得所有一级分组
    _groupId   int64
    // 补充信息
    _extraContext   string
}

// 初始化AlibabaIcbuProductGroupGetRequest对象
func NewAlibabaIcbuProductGroupGetRequest() *AlibabaIcbuProductGroupGetRequest{
    return &AlibabaIcbuProductGroupGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGroupGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.group.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 分组ID，传-1获得所有一级分组
func (r *AlibabaIcbuProductGroupGetRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuProductGroupGetRequest) GetGroupId() int64 {
    return r._groupId
}
// ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductGroupGetRequest) SetExtraContext(_extraContext string) error {
    r._extraContext = _extraContext
    r.Set("extra_context", _extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuProductGroupGetRequest) GetExtraContext() string {
    return r._extraContext
}
