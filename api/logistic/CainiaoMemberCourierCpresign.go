package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Cainiaomembercouriercpresign cp清理离职用户信息
// cainiao.member.courier.cpresign
//
// CP清理内部离职的用户信息
func Cainiaomembercouriercpresign(clt *core.SDKClient, req *logistic.CainiaomembercouriercpresignAPIRequest, session string) (*logistic.CainiaomembercouriercpresignAPIResponse, error) {
	var resp logistic.CainiaomembercouriercpresignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
