package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductBatchUpdateDisplayAPIRequest
商品批量上下架接口 API请求
alibaba.icbu.product.batch.update.display

给国际站的三方服务商提供批量上下架接口 */
type AlibabaIcbuProductBatchUpdateDisplayAPIRequest struct {
	model.Params
	// on表示上架，off表示下架
	_newDisplay string
	// 用逗号分隔的混淆id字符串
	_productIdList string
}

// New
