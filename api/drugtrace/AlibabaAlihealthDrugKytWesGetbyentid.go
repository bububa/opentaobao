package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetbyentid 根据企业主键查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyentid
//
// 根据企业主键查看企业详细信息
func AlibabaAlihealthDrugKytWesGetbyentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesGetbyentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
