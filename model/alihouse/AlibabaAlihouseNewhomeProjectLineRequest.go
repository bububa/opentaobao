package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
楼盘上下架 API请求
alibaba.alihouse.newhome.project.line

上下架楼盘
*/
type AlibabaAlihouseNewhomeProjectLineAPIRequest struct {
    model.Params
    // 外部id
    _outerId   string
    // 0-下架 1-上架
    _type   *model.File
}

// 初始化AlibabaAlihouseNewhomeProjectLineAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectLineRequest() *AlibabaAlihouseNewhomeProjectLineAPIRequest{
    return &AlibabaAlihouseNewhomeProjectLineAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.line"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 外部id
func (r *AlibabaAlihouseNewhomeProjectLineAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetOuterId() string {
    return r._outerId
}
// Type Setter
// 0-下架 1-上架
func (r *AlibabaAlihouseNewhomeProjectLineAPIRequest) SetType(_type *model.File) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlihouseNewhomeProjectLineAPIRequest) GetType() *model.File {
    return r._type
}
