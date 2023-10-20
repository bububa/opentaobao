package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmCustomerInfoGet 会员信息查询接口
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
func AlibabaWestcrmCustomerInfoGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmCustomerInfoGetAPIRequest, resp *westcrm.AlibabaWestcrmCustomerInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
