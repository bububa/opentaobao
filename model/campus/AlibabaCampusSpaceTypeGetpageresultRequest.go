package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询空间类别接口 APIRequest
alibaba.campus.space.type.getpageresult

分页查询空间类别接口
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getPageResult
*/
type AlibabaCampusSpaceTypeGetpageresultRequest struct {
    model.Params

    // 环境参数
    param0   *WorkBenchContext 

    // 查询参数
    param1   *SpaceTypeQuery 

}

func NewAlibabaCampusSpaceTypeGetpageresultRequest() *AlibabaCampusSpaceTypeGetpageresultRequest{
    return &AlibabaCampusSpaceTypeGetpageresultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceTypeGetpageresultRequest) GetApiMethodName() string {
    return "alibaba.campus.space.type.getpageresult"
}

func (r AlibabaCampusSpaceTypeGetpageresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceTypeGetpageresultRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusSpaceTypeGetpageresultRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusSpaceTypeGetpageresultRequest) SetParam1(param1 *SpaceTypeQuery) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusSpaceTypeGetpageresultRequest) GetParam1() *SpaceTypeQuery {
    return r.param1
}

