package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除主动预约开通条件 APIRequest
tmall.servicecenter.reservecond.delete

删除主动预约开通条件
*/
type TmallServicecenterReservecondDeleteRequest struct {
    model.Params

    // 主动预约条件删除
    reserveOpenConditionDelDto   *ReserveOpenConditionDelDto 

}

func NewTmallServicecenterReservecondDeleteRequest() *TmallServicecenterReservecondDeleteRequest{
    return &TmallServicecenterReservecondDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterReservecondDeleteRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.delete"
}

func (r TmallServicecenterReservecondDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterReservecondDeleteRequest) SetReserveOpenConditionDelDto(reserveOpenConditionDelDto *ReserveOpenConditionDelDto) error {
    r.reserveOpenConditionDelDto = reserveOpenConditionDelDto
    r.Set("reserve_open_condition_del_dto", reserveOpenConditionDelDto)
    return nil
}

func (r TmallServicecenterReservecondDeleteRequest) GetReserveOpenConditionDelDto() *ReserveOpenConditionDelDto {
    return r.reserveOpenConditionDelDto
}

