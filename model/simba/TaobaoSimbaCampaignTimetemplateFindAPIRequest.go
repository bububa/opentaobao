package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaigntimetemplatefindAPIRequest 获取分时折扣模板 API请求
// taobao.simba.campaign.timetemplate.find
//
// 批量得到智能推广推广计划下的推广组
type TaobaosimbacampaigntimetemplatefindAPIRequest struct {
	model.Params
}

// NewTaobaosimbacampaigntimetemplatefindRequest 初始化TaobaosimbacampaigntimetemplatefindAPIRequest对象
func NewTaobaosimbacampaigntimetemplatefindRequest() *TaobaosimbacampaigntimetemplatefindAPIRequest {
	return &TaobaosimbacampaigntimetemplatefindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaigntimetemplatefindAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.timetemplate.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaigntimetemplatefindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaigntimetemplatefindAPIRequest) GetRawParams() model.Params {
	return r.Params
}
