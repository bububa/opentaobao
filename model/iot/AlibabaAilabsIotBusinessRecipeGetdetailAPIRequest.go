package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest 获取食谱详情 API请求
// alibaba.ailabs.iot.business.recipe.getdetail
//
// 获取食谱详情接口，获取ISV自己的食谱详情数据
type AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest struct {
	model.Params
	// 行业食谱id
	_businessRecipeId int64
	// 开放账号id
	_openAccountId string
}

// NewAlibabaAilabsIotBusinessRecipeGetdetailRequest 初始化AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest对象
func NewAlibabaAilabsIotBusinessRecipeGetdetailRequest() *AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest {
	return &AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BusinessRecipeId Setter
// 行业食谱id
func (r *AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) SetBusinessRecipeId(_businessRecipeId int64) error {
	r._businessRecipeId = _businessRecipeId
	r.Set("business_recipe_id", _businessRecipeId)
	return nil
}

// Get BusinessRecipeId Getter
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetBusinessRecipeId() int64 {
	return r._businessRecipeId
}

// Set is OpenAccountId Setter
// 开放账号id
func (r *AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) SetOpenAccountId(_openAccountId string) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// Get OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest) GetOpenAccountId() string {
	return r._openAccountId
}
