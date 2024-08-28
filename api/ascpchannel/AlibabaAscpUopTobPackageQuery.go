package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopTobPackageQuery B2B包裹查询接口
// alibaba.ascp.uop.tob.package.query
//
// 供应链中台TOB包裹查询接口
func AlibabaAscpUopTobPackageQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopTobPackageQueryAPIRequest, resp *ascpchannel.AlibabaAscpUopTobPackageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
