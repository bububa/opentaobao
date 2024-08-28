package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalStoreInvisibleConsumeUpdate 门店无隐形消费签约
// alibaba.alihealth.dental.store.invisible.consume.update
//
// 门店无隐形消费签约
func AlibabaAlihealthDentalStoreInvisibleConsumeUpdate(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest, resp *alihealth2.AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
