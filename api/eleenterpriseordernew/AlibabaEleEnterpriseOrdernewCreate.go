package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewCreate 创建订单
// alibaba.ele.enterprise.ordernew.create
//
// 创建订单
func AlibabaEleEnterpriseOrdernewCreate(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCreateAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
