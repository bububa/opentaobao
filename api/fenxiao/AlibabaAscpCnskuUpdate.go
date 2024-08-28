package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuUpdate 供应链中台货品修改接口
// alibaba.ascp.cnsku.update
//
// 供应链中台货品修改接口
func AlibabaAscpCnskuUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuUpdateAPIRequest, resp *fenxiao.AlibabaAscpCnskuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
