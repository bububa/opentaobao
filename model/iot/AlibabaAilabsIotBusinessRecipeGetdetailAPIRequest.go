package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest
获取食谱详情 API请求
alibaba.ailabs.iot.business.recipe.getdetail

获取食谱详情接口，获取ISV自己的食谱详情数据 */
type AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest struct {
	model.Params
	// 行业食谱id
	_businessRecipeId int64
	// 开放账号id
	_openAccountId string
}

// New
