package alitripdivisions

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据父节点id查询下级行政区划数据 APIRequest
alitrip.platform.divisions.querybyparentid

根据行政区划id查询下一层级行政区划数据
*/
type AlitripPlatformDivisionsQuerybyparentidRequest struct {
    model.Params

    // 行政区划父id
    paramLong   int64 

}

func NewAlitripPlatformDivisionsQuerybyparentidRequest() *AlitripPlatformDivisionsQuerybyparentidRequest{
    return &AlitripPlatformDivisionsQuerybyparentidRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformDivisionsQuerybyparentidRequest) GetApiMethodName() string {
    return "alitrip.platform.divisions.querybyparentid"
}

func (r AlitripPlatformDivisionsQuerybyparentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformDivisionsQuerybyparentidRequest) SetParamLong(paramLong int64) error {
    r.paramLong = paramLong
    r.Set("param_long", paramLong)
    return nil
}

func (r AlitripPlatformDivisionsQuerybyparentidRequest) GetParamLong() int64 {
    return r.paramLong
}

