package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTraderateExplainAddAPIRequest
商城评价解释接口 API请求
taobao.traderate.explain.add

商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释） */
type TaobaoTraderateExplainAddAPIRequest struct {
	model.Params
	// 子订单ID
	_oid int64
	// 评价解释内容，最大长度：500个汉字
	_reply string
}

// New
