package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeCommonListCodeinfo 通用查询码接口
// alibaba.alihealth.drug.code.common.list.codeinfo
//
// 通用查询码接口
func AlibabaAlihealthDrugCodeCommonListCodeinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
