package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasestandpointqueryref 查询推送口径信息
// alibaba.legal.case.standpoint.queryref
//
// 查询推送口径信息
func Alibabalegalcasestandpointqueryref(clt *core.SDKClient, req *legalcase.AlibabalegalcasestandpointqueryrefAPIRequest, session string) (*legalcase.AlibabalegalcasestandpointqueryrefAPIResponse, error) {
	var resp legalcase.AlibabalegalcasestandpointqueryrefAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
