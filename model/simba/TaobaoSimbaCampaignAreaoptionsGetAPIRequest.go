package simba

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignAreaoptionsGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaoptionsGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.areaoptions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaoptionsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignAreaoptionsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoSimbaCampaignAreaoptionsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignAreaoptionsGetRequest()
	},
}

// GetTaobaoSimbaCampaignAreaoptionsGetRequest 从 sync.Pool 获取 TaobaoSimbaCampaignAreaoptionsGetAPIRequest
func GetTaobaoSimbaCampaignAreaoptionsGetAPIRequest() *TaobaoSimbaCampaignAreaoptionsGetAPIRequest {
	return poolTaobaoSimbaCampaignAreaoptionsGetAPIRequest.Get().(*TaobaoSimbaCampaignAreaoptionsGetAPIRequest)
}

// ReleaseTaobaoSimbaCampaignAreaoptionsGetAPIRequest 将 TaobaoSimbaCampaignAreaoptionsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignAreaoptionsGetAPIRequest(v *TaobaoSimbaCampaignAreaoptionsGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignAreaoptionsGetAPIRequest.Put(v)
}
