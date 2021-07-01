package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest
插入或更新食谱步骤 API请求
alibaba.ailabs.iot.business.recipestep.insertorupdate

插入或更新食谱步骤 */
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIRequest struct {
	model.Params
	// 食谱步骤开放参数
	_paramBusinessRecipeStepOpenParam *BusinessRecipeStepOpenParam
}

// New
