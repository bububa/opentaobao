package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanUpdateProductsAPIRequest 定向推广 按照id操作推广计划的产品，包括新增，删除和更新 API请求
// alibaba.scbp.target.ad.plan.update.products
//
// 定向推广 按照id操作推广计划的产品，包括新增，删除和更新
type AlibabaScbpTargetAdPlanUpdateProductsAPIRequest struct {
	model.Params
	// 系统生成
	_paramTopP4pModifyQuickCampaignProductDTO *TopP4pModifyQuickCampaignProductDto
}

// NewAlibabaScbpTargetAdPlanUpdateProductsRequest 初始化AlibabaScbpTargetAdPlanUpdateProductsAPIRequest对象
func NewAlibabaScbpTargetAdPlanUpdateProductsRequest() *AlibabaScbpTargetAdPlanUpdateProductsAPIRequest {
	return &AlibabaScbpTargetAdPlanUpdateProductsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) Reset() {
	r._paramTopP4pModifyQuickCampaignProductDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.update.products"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopP4pModifyQuickCampaignProductDTO is ParamTopP4pModifyQuickCampaignProductDTO Setter
// 系统生成
func (r *AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) SetParamTopP4pModifyQuickCampaignProductDTO(_paramTopP4pModifyQuickCampaignProductDTO *TopP4pModifyQuickCampaignProductDto) error {
	r._paramTopP4pModifyQuickCampaignProductDTO = _paramTopP4pModifyQuickCampaignProductDTO
	r.Set("param_top_p4p_modify_quick_campaign_product_d_t_o", _paramTopP4pModifyQuickCampaignProductDTO)
	return nil
}

// GetParamTopP4pModifyQuickCampaignProductDTO ParamTopP4pModifyQuickCampaignProductDTO Getter
func (r AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) GetParamTopP4pModifyQuickCampaignProductDTO() *TopP4pModifyQuickCampaignProductDto {
	return r._paramTopP4pModifyQuickCampaignProductDTO
}

var poolAlibabaScbpTargetAdPlanUpdateProductsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdPlanUpdateProductsRequest()
	},
}

// GetAlibabaScbpTargetAdPlanUpdateProductsRequest 从 sync.Pool 获取 AlibabaScbpTargetAdPlanUpdateProductsAPIRequest
func GetAlibabaScbpTargetAdPlanUpdateProductsAPIRequest() *AlibabaScbpTargetAdPlanUpdateProductsAPIRequest {
	return poolAlibabaScbpTargetAdPlanUpdateProductsAPIRequest.Get().(*AlibabaScbpTargetAdPlanUpdateProductsAPIRequest)
}

// ReleaseAlibabaScbpTargetAdPlanUpdateProductsAPIRequest 将 AlibabaScbpTargetAdPlanUpdateProductsAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanUpdateProductsAPIRequest(v *AlibabaScbpTargetAdPlanUpdateProductsAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanUpdateProductsAPIRequest.Put(v)
}
