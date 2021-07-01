package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAsrDataservicePromotionruleQueryAPIRequest
星巴克优惠规则查询 API请求
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则 */
type AlibabaAsrDataservicePromotionruleQueryAPIRequest struct {
	model.Params
	// 当前页
	_pageNo int64
	// 每页数量
	_pageSize int64
}

// New
