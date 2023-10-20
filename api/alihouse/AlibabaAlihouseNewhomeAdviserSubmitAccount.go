package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeAdviserSubmitAccount 顾问入职离职
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
func AlibabaAlihouseNewhomeAdviserSubmitAccount(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
