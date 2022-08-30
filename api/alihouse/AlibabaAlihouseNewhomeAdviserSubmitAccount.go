package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeAdviserSubmitAccount 顾问入职离职
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
func AlibabaAlihouseNewhomeAdviserSubmitAccount(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeAdviserSubmitAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
