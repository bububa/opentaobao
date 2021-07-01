package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemCalculateHscodeGetAPIRequest
算法获取hscode API请求
tmall.item.calculate.hscode.get

算法获取hscode */
type TmallItemCalculateHscodeGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
