package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasestandpointquerystandpoint 主动查询口径
// alibaba.legal.case.standpoint.querystandpoint
//
// 为法宝侧提供主动查询口径功能,有利于规范外部律师答辩口径.
func Alibabalegalcasestandpointquerystandpoint(clt *core.SDKClient, req *legalcase.AlibabalegalcasestandpointquerystandpointAPIRequest, session string) (*legalcase.AlibabalegalcasestandpointquerystandpointAPIResponse, error) {
	var resp legalcase.AlibabalegalcasestandpointquerystandpointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
