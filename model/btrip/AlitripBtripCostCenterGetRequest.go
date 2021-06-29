package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户费用归属 APIRequest
alitrip.btrip.cost.center.get

获取差旅申请用户的费用归属列表
*/
type AlitripBtripCostCenterGetRequest struct {
    model.Params

    // 企业id
    corpId   string 

    // 用户id
    userId   string 

}

func NewAlitripBtripCostCenterGetRequest() *AlitripBtripCostCenterGetRequest{
    return &AlitripBtripCostCenterGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCostCenterGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.get"
}

func (r AlitripBtripCostCenterGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCostCenterGetRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlitripBtripCostCenterGetRequest) GetCorpId() string {
    return r.corpId
}

func (r *AlitripBtripCostCenterGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlitripBtripCostCenterGetRequest) GetUserId() string {
    return r.userId
}

