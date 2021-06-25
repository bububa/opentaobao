package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动创建 APIRequest
taobao.jst.miniapp.crowd.create

小程序活动创建
*/
type TaobaoJstMiniappCrowdCreateRequest struct {
    model.Params

    // 活动开始时间，开始时间和结束时间不能超过1个月
    endDate   string 

    // 活动描述
    description   string 

    // 活动开始时间
    startDate   string 

    // 活动名称
    crowdName   string 

}

func NewTaobaoJstMiniappCrowdCreateRequest() *TaobaoJstMiniappCrowdCreateRequest{
    return &TaobaoJstMiniappCrowdCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstMiniappCrowdCreateRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.crowd.create"
}

func (r TaobaoJstMiniappCrowdCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstMiniappCrowdCreateRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoJstMiniappCrowdCreateRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoJstMiniappCrowdCreateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r TaobaoJstMiniappCrowdCreateRequest) GetDescription() string {
    return r.description
}

func (r *TaobaoJstMiniappCrowdCreateRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoJstMiniappCrowdCreateRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoJstMiniappCrowdCreateRequest) SetCrowdName(crowdName string) error {
    r.crowdName = crowdName
    r.Set("crowd_name", crowdName)
    return nil
}

func (r TaobaoJstMiniappCrowdCreateRequest) GetCrowdName() string {
    return r.crowdName
}

