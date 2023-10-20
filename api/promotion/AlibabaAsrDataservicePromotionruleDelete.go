package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabaasrdataservicepromotionruledelete 优惠规则删除
// alibaba.asr.dataservice.promotionrule.delete
//
// 删除优惠规则，例如星巴克删除优惠规则
func Alibabaasrdataservicepromotionruledelete(clt *core.SDKClient, req *promotion.AlibabaasrdataservicepromotionruledeleteAPIRequest, session string) (*promotion.AlibabaasrdataservicepromotionruledeleteAPIResponse, error) {
	var resp promotion.AlibabaasrdataservicepromotionruledeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
