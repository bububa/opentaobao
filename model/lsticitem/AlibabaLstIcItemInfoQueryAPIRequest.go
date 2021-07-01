package lsticitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstIcItemInfoQueryAPIRequest
商品信息查询 API请求
alibaba.lst.ic.item.info.query

查询商品信息 */
type AlibabaLstIcItemInfoQueryAPIRequest struct {
	model.Params
	// 零售通商品查询参数
	_query *LstItemListParam
}

// New
