package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasestandpointsavestandpoint 新增反馈口径
// alibaba.legal.case.standpoint.savestandpoint
//
// 新增反馈口径 ,从外部接受反馈的口径
func Alibabalegalcasestandpointsavestandpoint(clt *core.SDKClient, req *legalcase.AlibabalegalcasestandpointsavestandpointAPIRequest, session string) (*legalcase.AlibabalegalcasestandpointsavestandpointAPIResponse, error) {
	var resp legalcase.AlibabalegalcasestandpointsavestandpointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
