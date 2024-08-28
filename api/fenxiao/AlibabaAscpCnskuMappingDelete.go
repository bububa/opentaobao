package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuMappingDelete 货品关系解绑
// alibaba.ascp.cnsku.mapping.delete
//
// 货品关系解绑
func AlibabaAscpCnskuMappingDelete(ctx context.Context, clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuMappingDeleteAPIRequest, resp *fenxiao.AlibabaAscpCnskuMappingDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
