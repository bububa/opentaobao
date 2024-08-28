package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drug.kyt.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugKytGetbyrefentid(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetbyrefentidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
