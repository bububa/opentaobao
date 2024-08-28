package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeUserData 西安杨森同步用户行为接口
// alibaba.alihealth.drugcode.user.data
//
// 西安杨森同步用户行为接口
func AlibabaAlihealthDrugcodeUserData(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeUserDataAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeUserDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
