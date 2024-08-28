package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpSuborderEstcontimeModify 向前修改发货时效
// alibaba.ascp.suborder.estcontime.modify
//
// 向前修改发货时效
func AlibabaAscpSuborderEstcontimeModify(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaAscpSuborderEstcontimeModifyAPIRequest, resp *ascp.AlibabaAscpSuborderEstcontimeModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
