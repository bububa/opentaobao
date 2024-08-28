package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeAdviserSubmitAccount 顾问入职离职
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
func AlibabaAlihouseNewhomeAdviserSubmitAccount(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
