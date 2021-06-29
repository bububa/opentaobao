package alitripdivisions

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据父节点id查询下级行政区划数据 API请求
alitrip.platform.divisions.querybyparentid

根据行政区划id查询下一层级行政区划数据
*/
type AlitripPlatformDivisionsQuerybyparentidRequest struct {
    model.Params
    // 行政区划父id
    _paramLong   int64
}

// 初始化AlitripPlatformDivisionsQuerybyparentidRequest对象
func NewAlitripPlatformDivisionsQuerybyparentidRequest() *AlitripPlatformDivisionsQuerybyparentidRequest{
    return &AlitripPlatformDivisionsQuerybyparentidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformDivisionsQuerybyparentidRequest) GetApiMethodName() string {
    return "alitrip.platform.divisions.querybyparentid"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformDivisionsQuerybyparentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamLong Setter
// 行政区划父id
func (r *AlitripPlatformDivisionsQuerybyparentidRequest) SetParamLong(_paramLong int64) error {
    r._paramLong = _paramLong
    r.Set("param_long", _paramLong)
    return nil
}

// ParamLong Getter
func (r AlitripPlatformDivisionsQuerybyparentidRequest) GetParamLong() int64 {
    return r._paramLong
}
