package newretail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/newretail"
)

// AlibabaItApAddressGet getApAddressByMacNew
// alibaba.it.ap.address.get
//
// 根据ap 的mac地址查询ap的结构化位置信息
func AlibabaItApAddressGet(ctx context.Context, clt *core.SDKClient, req *newretail.AlibabaItApAddressGetAPIRequest, resp *newretail.AlibabaItApAddressGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
