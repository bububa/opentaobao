package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantEnterpriseEntry 机构入驻
// alibaba.alihouse.merchant.enterprise.entry
//
// 机构入驻
func AlibabaAlihouseMerchantEnterpriseEntry(clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIRequest, session string) (*alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
