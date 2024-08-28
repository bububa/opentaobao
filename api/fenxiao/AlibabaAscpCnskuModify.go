package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuModify 供应链中台货品修改接口
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
func AlibabaAscpCnskuModify(ctx context.Context, clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuModifyAPIRequest, resp *fenxiao.AlibabaAscpCnskuModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
