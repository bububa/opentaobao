package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewGetstatus 订单状态查询接口
// alibaba.ele.enterprise.ordernew.getstatus
//
// 订单状态查询接口
func AlibabaEleEnterpriseOrdernewGetstatus(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetstatusAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
