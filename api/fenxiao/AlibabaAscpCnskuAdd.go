package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuAdd 货品创建
// alibaba.ascp.cnsku.add
//
// 供应链中台货品创建接口
func AlibabaAscpCnskuAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuAddAPIRequest, resp *fenxiao.AlibabaAscpCnskuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
