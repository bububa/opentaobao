package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAsrDataservicePromotionruleWriteAPIRequest
业务优惠规则写入 API请求
alibaba.asr.dataservice.promotionrule.write

星巴克优惠规则写入 */
type AlibabaAsrDataservicePromotionruleWriteAPIRequest struct {
	model.Params
	// 入参对象
	_poskeyPromotionRuleDto *PosKeyPromotionRuleDto
}

// New
