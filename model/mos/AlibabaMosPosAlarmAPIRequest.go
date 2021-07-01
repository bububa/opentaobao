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
type AlibabaMosPosAlarmAPIRequest struct {
    model.Params
    // 参数
    _param0   *PosLogDto
}

// 初始化AlibabaMosPosAlarmAPIRequest对象
func NewAlibabaMosPosAlarmRequest() *AlibabaMosPosAlarmAPIRequest{
    return &AlibabaMosPosAlarmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosPosAlarmAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.pos.alarm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosPosAlarmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *AlibabaMosPosAlarmAPIRequest) SetParam0(_param0 *PosLogDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaMosPosAlarmAPIRequest) GetParam0() *PosLogDto {
    return r._param0
}
