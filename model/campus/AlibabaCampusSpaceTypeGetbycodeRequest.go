package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据类别编码查询类别 API请求
alibaba.campus.space.type.getbycode

根据类别编码查询类别
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getByCode
*/
type AlibabaCampusSpaceTypeGetbycodeRequest struct {
    model.Params
    // 查询条件封装
    _param0   *WorkBenchContext
    // 空间类别编码
    _typeCode   string
}

// 初始化AlibabaCampusSpaceTypeGetbycodeRequest对象
func NewAlibabaCampusSpaceTypeGetbycodeRequest() *AlibabaCampusSpaceTypeGetbycodeRequest{
    return &AlibabaCampusSpaceTypeGetbycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceTypeGetbycodeRequest) GetApiMethodName() string {
    return "alibaba.campus.space.type.getbycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceTypeGetbycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceTypeGetbycodeRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusSpaceTypeGetbycodeRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// TypeCode Setter
// 空间类别编码
func (r *AlibabaCampusSpaceTypeGetbycodeRequest) SetTypeCode(_typeCode string) error {
    r._typeCode = _typeCode
    r.Set("type_code", _typeCode)
    return nil
}

// TypeCode Getter
func (r AlibabaCampusSpaceTypeGetbycodeRequest) GetTypeCode() string {
    return r._typeCode
}
