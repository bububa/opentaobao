package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotbusinessrecipegetdetailAPIRequest 获取食谱详情 API请求
// alibaba.ailabs.iot.business.recipe.getdetail
//
// 获取食谱详情接口，获取ISV自己的食谱详情数据
type AlibabaailabsiotbusinessrecipegetdetailAPIRequest struct {
	model.Params
	// 开放账号id
	_openAccountId string
	// 行业食谱id
	_businessRecipeId int64
}

// NewAlibabaailabsiotbusinessrecipegetdetailRequest 初始化AlibabaailabsiotbusinessrecipegetdetailAPIRequest对象
func NewAlibabaailabsiotbusinessrecipegetdetailRequest() *AlibabaailabsiotbusinessrecipegetdetailAPIRequest {
	return &AlibabaailabsiotbusinessrecipegetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotbusinessrecipegetdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotbusinessrecipegetdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotbusinessrecipegetdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAccountId is OpenAccountId Setter
// 开放账号id
func (r *AlibabaailabsiotbusinessrecipegetdetailAPIRequest) SetOpenAccountId(_openAccountId string) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// GetOpenAccountId OpenAccountId Getter
func (r AlibabaailabsiotbusinessrecipegetdetailAPIRequest) GetOpenAccountId() string {
	return r._openAccountId
}

// SetBusinessRecipeId is BusinessRecipeId Setter
// 行业食谱id
func (r *AlibabaailabsiotbusinessrecipegetdetailAPIRequest) SetBusinessRecipeId(_businessRecipeId int64) error {
	r._businessRecipeId = _businessRecipeId
	r.Set("business_recipe_id", _businessRecipeId)
	return nil
}

// GetBusinessRecipeId BusinessRecipeId Getter
func (r AlibabaailabsiotbusinessrecipegetdetailAPIRequest) GetBusinessRecipeId() int64 {
	return r._businessRecipeId
}
