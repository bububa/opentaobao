package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
pos故障报警 APIRequest
alibaba.mos.pos.alarm

故障报警
*/
type AlibabaMosPosAlarmRequest struct {
    model.Params

    // 参数
    param0   *PosLogDto 

}

func NewAlibabaMosPosAlarmRequest() *AlibabaMosPosAlarmRequest{
    return &AlibabaMosPosAlarmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosPosAlarmRequest) GetApiMethodName() string {
    return "alibaba.mos.pos.alarm"
}

func (r AlibabaMosPosAlarmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosPosAlarmRequest) SetParam0(param0 *PosLogDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaMosPosAlarmRequest) GetParam0() *PosLogDto {
    return r.param0
}

