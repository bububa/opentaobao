package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
pos故障报警 API请求
alibaba.mos.pos.alarm

故障报警
*/
type AlibabaMosPosAlarmRequest struct {
    model.Params
    // 参数
    param0   *PosLogDto
}

// 初始化AlibabaMosPosAlarmRequest对象
func NewAlibabaMosPosAlarmRequest() *AlibabaMosPosAlarmRequest{
    return &AlibabaMosPosAlarmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosPosAlarmRequest) GetApiMethodName() string {
    return "alibaba.mos.pos.alarm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosPosAlarmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *AlibabaMosPosAlarmRequest) SetParam0(param0 *PosLogDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaMosPosAlarmRequest) GetParam0() *PosLogDto {
    return r.param0
}
