package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得推广计划的可设置投放地域列表 APIRequest
taobao.simba.campaign.areaoptions.get

取得推广计划的可设置投放地域列表
*/
type TaobaoSimbaCampaignAreaoptionsGetRequest struct {
    model.Params

}

func NewTaobaoSimbaCampaignAreaoptionsGetRequest() *TaobaoSimbaCampaignAreaoptionsGetRequest{
    return &TaobaoSimbaCampaignAreaoptionsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignAreaoptionsGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.areaoptions.get"
}

func (r TaobaoSimbaCampaignAreaoptionsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


