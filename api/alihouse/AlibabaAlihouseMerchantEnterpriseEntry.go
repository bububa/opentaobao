package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantEnterpriseEntry 机构入驻
// alibaba.alihouse.merchant.enterprise.entry
//
// 机构入驻
func AlibabaAlihouseMerchantEnterpriseEntry(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIRequest, resp *alihouse.AlibabaAlihouseMerchantEnterpriseEntryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
