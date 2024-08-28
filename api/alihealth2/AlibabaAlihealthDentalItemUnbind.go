package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalItemUnbind ISV解绑商品
// alibaba.alihealth.dental.item.unbind
//
// ISV解绑商品
func AlibabaAlihealthDentalItemUnbind(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemUnbindAPIRequest, resp *alihealth2.AlibabaAlihealthDentalItemUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
