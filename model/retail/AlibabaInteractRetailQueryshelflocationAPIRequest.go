package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractRetailQueryshelflocationAPIRequest
查询货架和位置数据 API请求
alibaba.interact.retail.queryshelflocation

查询货架和位置数据 */
type AlibabaInteractRetailQueryshelflocationAPIRequest struct {
	model.Params
	// 门店code
	_param0 string
}

// New
