package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugKytWesGetbyrefentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
