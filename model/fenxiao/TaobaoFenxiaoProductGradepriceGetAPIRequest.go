package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductGradepriceGetAPIRequest
等级折扣查询 API请求
taobao.fenxiao.product.gradeprice.get

等级折扣查询 */
type TaobaoFenxiaoProductGradepriceGetAPIRequest struct {
	model.Params
	// 产品id
	_productId int64
	// skuId
	_skuId int64
	// 经、代销模式（1：代销、2：经销）
	_tradeType int64
}

// New
