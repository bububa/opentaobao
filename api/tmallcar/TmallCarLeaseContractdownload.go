package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseContractdownload 天猫开新车租后合同下载
// tmall.car.lease.contractdownload
//
// 天猫开新车租后合同下载
func TmallCarLeaseContractdownload(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeaseContractdownloadAPIRequest, resp *tmallcar.TmallCarLeaseContractdownloadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
