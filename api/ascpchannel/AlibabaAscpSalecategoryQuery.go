package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpSalecategoryQuery 货品品类查询
// alibaba.ascp.salecategory.query
//
// 根据货品ID查询对应销售品类ID
func AlibabaAscpSalecategoryQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpSalecategoryQueryAPIRequest, resp *ascpchannel.AlibabaAscpSalecategoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
