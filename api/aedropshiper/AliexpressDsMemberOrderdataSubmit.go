package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsMemberOrderdataSubmit dropshipper数据回流
// aliexpress.ds.member.orderdata.submit
//
// dropshipper数据回流
func AliexpressDsMemberOrderdataSubmit(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressDsMemberOrderdataSubmitAPIRequest, resp *aedropshiper.AliexpressDsMemberOrderdataSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
