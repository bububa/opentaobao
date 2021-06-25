package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建主动预约开通条件 APIRequest
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondCreateRequest struct {
    model.Params

    // 主动预约开通条件
    rocList   []ReserveOpenConditionDTO 

}

func NewTmallServicecenterReservecondCreateRequest() *TmallServicecenterReservecondCreateRequest{
    return &TmallServicecenterReservecondCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterReservecondCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.create"
}

func (r TmallServicecenterReservecondCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterReservecondCreateRequest) SetRocList(rocList []ReserveOpenConditionDTO) error {
    r.rocList = rocList
    r.Set("roc_list", rocList)
    return nil
}

func (r TmallServicecenterReservecondCreateRequest) GetRocList() []ReserveOpenConditionDTO {
    return r.rocList
}

