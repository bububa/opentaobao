package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabaasrdataservicepromotionrulewrite 业务优惠规则写入
// alibaba.asr.dataservice.promotionrule.write
//
// 星巴克优惠规则写入
func Alibabaasrdataservicepromotionrulewrite(clt *core.SDKClient, req *promotion.AlibabaasrdataservicepromotionrulewriteAPIRequest, session string) (*promotion.AlibabaasrdataservicepromotionrulewriteAPIResponse, error) {
	var resp promotion.AlibabaasrdataservicepromotionrulewriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
