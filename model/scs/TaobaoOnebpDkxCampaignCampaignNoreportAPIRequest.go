package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest 获取场景计划的非报表数据 API请求
// taobao.onebp.dkx.campaign.campaign.noreport
//
// 获取场景计划的非报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 计划查询参数
	_campaignQuery *CampaignQueryTopDto
}

// NewTaobaoOnebpDkxCampaignCampaignNoreportRequest 初始化TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest对象
func NewTaobaoOnebpDkxCampaignCampaignNoreportRequest() *TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest {
	return &TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._campaignQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.campaign.campaign.noreport"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetCampaignQuery is CampaignQuery Setter
// 计划查询参数
func (r *TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryTopDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) GetCampaignQuery() *CampaignQueryTopDto {
	return r._campaignQuery
}

var poolTaobaoOnebpDkxCampaignCampaignNoreportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxCampaignCampaignNoreportRequest()
	},
}

// GetTaobaoOnebpDkxCampaignCampaignNoreportRequest 从 sync.Pool 获取 TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest
func GetTaobaoOnebpDkxCampaignCampaignNoreportAPIRequest() *TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest {
	return poolTaobaoOnebpDkxCampaignCampaignNoreportAPIRequest.Get().(*TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest)
}

// ReleaseTaobaoOnebpDkxCampaignCampaignNoreportAPIRequest 将 TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxCampaignCampaignNoreportAPIRequest(v *TaobaoOnebpDkxCampaignCampaignNoreportAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxCampaignCampaignNoreportAPIRequest.Put(v)
}
