package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeApartmentLine 公寓上下架
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
func AlibabaAlihouseNewhomeApartmentLine(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeApartmentLineAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeApartmentLineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
