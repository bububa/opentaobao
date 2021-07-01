package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpShowcaseAddproductAPIRequest
批量添加橱窗商品 API请求
alibaba.scbp.showcase.addproduct

批量添加商品到橱窗 */
type AlibabaScbpShowcaseAddproductAPIRequest struct {
	model.Params
	// 需要添加的产品ids
	_productIdList []int64
}

// New
