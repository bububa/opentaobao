package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片银行分组信息获取 API请求
alibaba.icbu.photobank.group.list

图片银行分组信息获取
*/
type AlibabaIcbuPhotobankGroupListRequest struct {
    model.Params
    // 补充信息
    _extraContext   string
    // 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
    _id   int64
}

// 初始化AlibabaIcbuPhotobankGroupListRequest对象
func NewAlibabaIcbuPhotobankGroupListRequest() *AlibabaIcbuPhotobankGroupListRequest{
    return &AlibabaIcbuPhotobankGroupListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankGroupListRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.group.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankGroupListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuPhotobankGroupListRequest) SetExtraContext(_extraContext string) error {
    r._extraContext = _extraContext
    r.Set("extra_context", _extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuPhotobankGroupListRequest) GetExtraContext() string {
    return r._extraContext
}
// Id Setter
// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
func (r *AlibabaIcbuPhotobankGroupListRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaIcbuPhotobankGroupListRequest) GetId() int64 {
    return r._id
}
