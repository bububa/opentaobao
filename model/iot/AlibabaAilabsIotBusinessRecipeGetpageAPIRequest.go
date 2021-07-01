package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotBusinessRecipeGetpageAPIRequest
分页查询食谱 API请求
alibaba.ailabs.iot.business.recipe.getpage

分页查询食谱数据 */
type AlibabaAilabsIotBusinessRecipeGetpageAPIRequest struct {
	model.Params
	// 开放账号id
	_openAccountId string
	// 分页页码
	_pageNum int64
	// 分页大小
	_pageSize int64
}

// New
