package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaAsrDataservicePromotionruleWrite 业务优惠规则写入
// alibaba.asr.dataservice.promotionrule.write
//
// 星巴克优惠规则写入
func AlibabaAsrDataservicePromotionruleWrite(clt *core.SDKClient, req *promotion.AlibabaAsrDataservicePromotionruleWriteAPIRequest, session string) (*promotion.AlibabaAsrDataservicePromotionruleWriteAPIResponse, error) {
	var resp promotion.AlibabaAsrDataservicePromotionruleWriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
