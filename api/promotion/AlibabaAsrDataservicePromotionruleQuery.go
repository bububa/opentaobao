package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabaasrdataservicepromotionrulequery 星巴克优惠规则查询
// alibaba.asr.dataservice.promotionrule.query
//
// 查询优惠规则，例如星巴克查询优惠规则
func Alibabaasrdataservicepromotionrulequery(clt *core.SDKClient, req *promotion.AlibabaasrdataservicepromotionrulequeryAPIRequest, session string) (*promotion.AlibabaasrdataservicepromotionrulequeryAPIResponse, error) {
	var resp promotion.AlibabaasrdataservicepromotionrulequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
