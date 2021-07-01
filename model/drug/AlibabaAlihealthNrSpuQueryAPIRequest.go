package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrSpuQueryAPIRequest
获取标品库标品信息 API请求
alibaba.alihealth.nr.spu.query

提供给ERP使用的，获取健康标品库信息 */
type AlibabaAlihealthNrSpuQueryAPIRequest struct {
	model.Params
	// 标品查询条件
	_query *TopAlihealthSpuQuery
	// 查询选择器
	_options *TopAlihealthSpuQueryOptions
}

// New
