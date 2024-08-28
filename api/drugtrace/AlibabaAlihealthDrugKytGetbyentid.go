package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetbyentid 根据企业主键查看企业详细信息
// alibaba.alihealth.drug.kyt.getbyentid
//
// 根据企业主键查看企业详细信息
func AlibabaAlihealthDrugKytGetbyentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetbyentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetbyentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
