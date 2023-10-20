package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignTimetemplateFindAPIRequest 获取分时折扣模板 API请求
// taobao.simba.campaign.timetemplate.find
//
// 批量得到智能推广推广计划下的推广组
type TaobaoSimbaCampaignTimetemplateFindAPIRequest struct {
	model.Params
}

// NewTaobaoSimbaCampaignTimetemplateFindRequest 初始化TaobaoSimbaCampaignTimetemplateFindAPIRequest对象
func NewTaobaoSimbaCampaignTimetemplateFindRequest() *TaobaoSimbaCampaignTimetemplateFindAPIRequest {
	return &TaobaoSimbaCampaignTimetemplateFindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignTimetemplateFindAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.timetemplate.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignTimetemplateFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignTimetemplateFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}
