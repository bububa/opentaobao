package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryIcpQueryLbx icp订单号查询lbx订单号
// alibaba.ascp.industry.icp.query.lbx
//
// 根据icp订单号查询lbx订单号
func AlibabaAscpIndustryIcpQueryLbx(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryIcpQueryLbxAPIRequest, resp *ascpchannel.AlibabaAscpIndustryIcpQueryLbxAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
