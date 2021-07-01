package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest
插入和更新食谱 API请求
alibaba.ailabs.iot.business.recipe.insertorupdate

插入和更新食谱，将isv的食谱添加到云端进行存储 */
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest struct {
	model.Params
	// 行业食谱开放参数
	_paramBusinessRecipeOpenParam *BusinessRecipeOpenParam
}

// New
