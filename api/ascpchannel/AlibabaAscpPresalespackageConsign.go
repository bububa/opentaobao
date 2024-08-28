package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpPresalespackageConsign 预售预包尾款推单发货
// alibaba.ascp.presalespackage.consign
//
// 预售预包尾款发货后推单处理
func AlibabaAscpPresalespackageConsign(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpPresalespackageConsignAPIRequest, resp *ascpchannel.AlibabaAscpPresalespackageConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
