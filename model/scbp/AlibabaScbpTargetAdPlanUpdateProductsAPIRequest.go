package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanupdateproductsAPIRequest 定向推广 按照id操作推广计划的产品，包括新增，删除和更新 API请求
// alibaba.scbp.target.ad.plan.update.products
//
// 定向推广 按照id操作推广计划的产品，包括新增，删除和更新
type AlibabascbptargetadplanupdateproductsAPIRequest struct {
	model.Params
	// 系统生成
	_paramTopP4pModifyQuickCampaignProductDTO *TopP4pModifyQuickCampaignProductDto
}

// NewAlibabascbptargetadplanupdateproductsRequest 初始化AlibabascbptargetadplanupdateproductsAPIRequest对象
func NewAlibabascbptargetadplanupdateproductsRequest() *AlibabascbptargetadplanupdateproductsAPIRequest {
	return &AlibabascbptargetadplanupdateproductsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplanupdateproductsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.update.products"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplanupdateproductsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplanupdateproductsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopP4pModifyQuickCampaignProductDTO is ParamTopP4pModifyQuickCampaignProductDTO Setter
// 系统生成
func (r *AlibabascbptargetadplanupdateproductsAPIRequest) SetParamTopP4pModifyQuickCampaignProductDTO(_paramTopP4pModifyQuickCampaignProductDTO *TopP4pModifyQuickCampaignProductDto) error {
	r._paramTopP4pModifyQuickCampaignProductDTO = _paramTopP4pModifyQuickCampaignProductDTO
	r.Set("param_top_p4p_modify_quick_campaign_product_d_t_o", _paramTopP4pModifyQuickCampaignProductDTO)
	return nil
}

// GetParamTopP4pModifyQuickCampaignProductDTO ParamTopP4pModifyQuickCampaignProductDTO Getter
func (r AlibabascbptargetadplanupdateproductsAPIRequest) GetParamTopP4pModifyQuickCampaignProductDTO() *TopP4pModifyQuickCampaignProductDto {
	return r._paramTopP4pModifyQuickCampaignProductDTO
}
