package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalItemBind ISV绑定外部门店id和外部商品id
// alibaba.alihealth.dental.item.bind
//
// ISV绑定外部门店id和外部商品id
func AlibabaAlihealthDentalItemBind(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemBindAPIRequest, resp *alihealth2.AlibabaAlihealthDentalItemBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
