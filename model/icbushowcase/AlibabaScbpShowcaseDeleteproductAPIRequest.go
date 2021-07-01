package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpShowcaseDeleteproductAPIRequest
批量删除橱窗商品 API请求
alibaba.scbp.showcase.deleteproduct

批量删除橱窗商品 */
type AlibabaScbpShowcaseDeleteproductAPIRequest struct {
	model.Params
	// 橱窗idList
	_windowIdList []int64
}

// New
