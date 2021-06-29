package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分时折扣模板 API请求
taobao.simba.campaign.timetemplate.find

批量得到智能推广推广计划下的推广组
*/
type TaobaoSimbaCampaignTimetemplateFindRequest struct {
    model.Params
}

// 初始化TaobaoSimbaCampaignTimetemplateFindRequest对象
func NewTaobaoSimbaCampaignTimetemplateFindRequest() *TaobaoSimbaCampaignTimetemplateFindRequest{
    return &TaobaoSimbaCampaignTimetemplateFindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignTimetemplateFindRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.timetemplate.find"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignTimetemplateFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
