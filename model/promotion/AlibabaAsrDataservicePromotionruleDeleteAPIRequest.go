package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAsrDataservicePromotionruleDeleteAPIRequest
优惠规则删除 API请求
alibaba.asr.dataservice.promotionrule.delete

删除优惠规则，例如星巴克删除优惠规则 */
type AlibabaAsrDataservicePromotionruleDeleteAPIRequest struct {
	model.Params
	// poskey
	_posKey int64
}

// New
