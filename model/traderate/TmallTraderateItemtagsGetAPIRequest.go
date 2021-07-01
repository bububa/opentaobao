package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraderateItemtagsGetAPIRequest
通过商品ID获取标签列表 API请求
tmall.traderate.itemtags.get

通过商品ID获取标签详细信息 */
type TmallTraderateItemtagsGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
