package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-楼层接口deleteFloor API请求
alibaba.damai.mev.open.deletefloor

deleteFloor
*/
type AlibabaDamaiMevOpenDeletefloorRequest struct {
    model.Params
    // 入参deleteFloorParam
    _deleteFloorParam   *FloorIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletefloorRequest对象
func NewAlibabaDamaiMevOpenDeletefloorRequest() *AlibabaDamaiMevOpenDeletefloorRequest{
    return &AlibabaDamaiMevOpenDeletefloorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletefloorRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletefloor"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletefloorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteFloorParam Setter
// 入参deleteFloorParam
func (r *AlibabaDamaiMevOpenDeletefloorRequest) SetDeleteFloorParam(_deleteFloorParam *FloorIdOpenParam) error {
    r._deleteFloorParam = _deleteFloorParam
    r.Set("delete_floor_param", _deleteFloorParam)
    return nil
}

// DeleteFloorParam Getter
func (r AlibabaDamaiMevOpenDeletefloorRequest) GetDeleteFloorParam() *FloorIdOpenParam {
    return r._deleteFloorParam
}
