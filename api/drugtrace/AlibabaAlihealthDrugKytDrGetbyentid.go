package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrGetbyentid 多融通过企业ID得到一个企业的详细信息
// alibaba.alihealth.drug.kyt.dr.getbyentid
//
// 根据企业主键查看企业详细信息
func AlibabaAlihealthDrugKytDrGetbyentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetbyentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrGetbyentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
