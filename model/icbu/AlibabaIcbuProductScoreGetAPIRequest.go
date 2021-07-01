package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductScoreGetAPIRequest
产品质量分查询 API请求
alibaba.icbu.product.score.get

产品质量分查询 */
type AlibabaIcbuProductScoreGetAPIRequest struct {
	model.Params
	// 混淆后的商品ID
	_productId string
}

// New
