package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpShowcaseUpdateproductAPIRequest
替换橱窗商品 API请求
alibaba.scbp.showcase.updateproduct

替换橱窗商品 */
type AlibabaScbpShowcaseUpdateproductAPIRequest struct {
	model.Params
	// 橱窗id
	_windowId int64
	// 新的商品id
	_newProductId int64
}

// New
