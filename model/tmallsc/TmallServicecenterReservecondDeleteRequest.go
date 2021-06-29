package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除主动预约开通条件 API请求
tmall.servicecenter.reservecond.delete

删除主动预约开通条件
*/
type TmallServicecenterReservecondDeleteRequest struct {
    model.Params
    // 主动预约条件删除
    _reserveOpenConditionDelDto   *ReserveOpenConditionDelDTO
}

// 初始化TmallServicecenterReservecondDeleteRequest对象
func NewTmallServicecenterReservecondDeleteRequest() *TmallServicecenterReservecondDeleteRequest{
    return &TmallServicecenterReservecondDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondDeleteRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.delete"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReserveOpenConditionDelDto Setter
// 主动预约条件删除
func (r *TmallServicecenterReservecondDeleteRequest) SetReserveOpenConditionDelDto(_reserveOpenConditionDelDto *ReserveOpenConditionDelDTO) error {
    r._reserveOpenConditionDelDto = _reserveOpenConditionDelDto
    r.Set("reserve_open_condition_del_dto", _reserveOpenConditionDelDto)
    return nil
}

// ReserveOpenConditionDelDto Getter
func (r TmallServicecenterReservecondDeleteRequest) GetReserveOpenConditionDelDto() *ReserveOpenConditionDelDTO {
    return r._reserveOpenConditionDelDto
}
