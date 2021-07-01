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
type TaobaoSimbaCampaignTimetemplateFindAPIRequest struct {
    model.Params
}

// 初始化TaobaoSimbaCampaignTimetemplateFindAPIRequest对象
func NewTaobaoSimbaCampaignTimetemplateFindRequest() *TaobaoSimbaCampaignTimetemplateFindAPIRequest{
    return &TaobaoSimbaCampaignTimetemplateFindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignTimetemplateFindAPIRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.timetemplate.find"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignTimetemplateFindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
