package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemGetAPIRequest
根据id查询商品 API请求
taobao.scitem.get

根据id查询商品 */
type TaobaoScitemGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
