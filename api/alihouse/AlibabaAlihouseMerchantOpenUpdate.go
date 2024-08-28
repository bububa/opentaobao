package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantOpenUpdate 非融合店进件升级成融合店
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
func AlibabaAlihouseMerchantOpenUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantOpenUpdateAPIRequest, resp *alihouse.AlibabaAlihouseMerchantOpenUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
