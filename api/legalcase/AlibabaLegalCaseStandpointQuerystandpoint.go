package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseStandpointQuerystandpoint 主动查询口径
// alibaba.legal.case.standpoint.querystandpoint
//
// 为法宝侧提供主动查询口径功能,有利于规范外部律师答辩口径.
func AlibabaLegalCaseStandpointQuerystandpoint(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointQuerystandpointAPIRequest, resp *legalcase.AlibabaLegalCaseStandpointQuerystandpointAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
