package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantEnterpriseEntry 机构入驻
// alibaba.alihouse.merchant.enterprise.entry
//
// 机构入驻
func AlibabaAlihouseMerchantEnterpriseEntry(clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIRequest, resp *alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
