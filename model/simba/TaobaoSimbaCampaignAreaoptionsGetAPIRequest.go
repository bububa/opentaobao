package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaoptionsGetAPIRequest 取得推广计划的可设置投放地域列表 API请求
// taobao.simba.campaign.areaoptions.get
//
// 取得推广计划的可设置投放地域列表
type TaobaoSimbaCampaignAreaoptionsGetAPIRequest struct {
	model.Params
}

// NewTaobaoSimbaCampaignAreaoptionsGetRequest 初始化TaobaoSimbaCampaignAreaoptionsGetAPIRequest对象
func NewTaobaoSimbaCampaignAreaoptionsGetRequest() *TaobaoSimbaCampaignAreaoptionsGetAPIRequest {
	return &TaobaoSimbaCampaignAreaoptionsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaoptionsGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.areaoptions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaoptionsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
