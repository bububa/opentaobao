package newretail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/newretail"
)

// AlibabaItApAddressSet setApAddressNew
// alibaba.it.ap.address.set
//
// 该接口可为ISV系统提供 ap位置信息维护的功能
func AlibabaItApAddressSet(ctx context.Context, clt *core.SDKClient, req *newretail.AlibabaItApAddressSetAPIRequest, resp *newretail.AlibabaItApAddressSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
