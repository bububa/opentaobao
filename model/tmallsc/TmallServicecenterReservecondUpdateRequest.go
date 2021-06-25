package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
主动预约条件更新 APIRequest
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondUpdateRequest struct {
    model.Params

    // 主动预约开通条件
    rocList   []ReserveOpenConditionDTO 

}

func NewTmallServicecenterReservecondUpdateRequest() *TmallServicecenterReservecondUpdateRequest{
    return &TmallServicecenterReservecondUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterReservecondUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.update"
}

func (r TmallServicecenterReservecondUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterReservecondUpdateRequest) SetRocList(rocList []ReserveOpenConditionDTO) error {
    r.rocList = rocList
    r.Set("roc_list", rocList)
    return nil
}

func (r TmallServicecenterReservecondUpdateRequest) GetRocList() []ReserveOpenConditionDTO {
    return r.rocList
}

