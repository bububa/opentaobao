package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaAsrDataservicePromotionruleWrite 业务优惠规则写入
// alibaba.asr.dataservice.promotionrule.write
//
// 星巴克优惠规则写入
func AlibabaAsrDataservicePromotionruleWrite(clt *core.SDKClient, req *promotion.AlibabaAsrDataservicePromotionruleWriteAPIRequest, resp *promotion.AlibabaAsrDataservicePromotionruleWriteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
