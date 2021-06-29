package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得推广计划的可设置投放地域列表 API请求
taobao.simba.campaign.areaoptions.get

取得推广计划的可设置投放地域列表
*/
type TaobaoSimbaCampaignAreaoptionsGetRequest struct {
    model.Params
}

// 初始化TaobaoSimbaCampaignAreaoptionsGetRequest对象
func NewTaobaoSimbaCampaignAreaoptionsGetRequest() *TaobaoSimbaCampaignAreaoptionsGetRequest{
    return &TaobaoSimbaCampaignAreaoptionsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaoptionsGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.areaoptions.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaoptionsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
