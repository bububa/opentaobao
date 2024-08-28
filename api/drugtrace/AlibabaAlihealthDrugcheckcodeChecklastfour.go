package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcheckcodeChecklastfour 校验追溯码的后4位是否正确
// alibaba.alihealth.drugcheckcode.checklastfour
//
// 校验追溯码的后4位是否正确
func AlibabaAlihealthDrugcheckcodeChecklastfour(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcheckcodeChecklastfourAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
