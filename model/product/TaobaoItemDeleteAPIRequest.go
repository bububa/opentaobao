package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemDeleteAPIRequest
删除单条商品 API请求
taobao.item.delete

删除单条商品 */
type TaobaoItemDeleteAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// New
