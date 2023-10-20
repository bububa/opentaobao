package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseStandpointQueryref 查询推送口径信息
// alibaba.legal.case.standpoint.queryref
//
// 查询推送口径信息
func AlibabaLegalCaseStandpointQueryref(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointQueryrefAPIRequest, session string) (*legalcase.AlibabaLegalCaseStandpointQueryrefAPIResponse, error) {
	var resp legalcase.AlibabaLegalCaseStandpointQueryrefAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
