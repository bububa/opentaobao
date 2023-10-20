package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewGetrefundinfo 退单和申诉
// alibaba.ele.enterprise.ordernew.getrefundinfo
//
// 退单和申诉
func AlibabaEleEnterpriseOrdernewGetrefundinfo(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGetrefundinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
