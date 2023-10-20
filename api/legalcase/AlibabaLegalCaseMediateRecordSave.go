package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasemediaterecordsave 新增调解结果
// alibaba.legal.case.mediate.record.save
//
// 增加调解沟通记录
func Alibabalegalcasemediaterecordsave(clt *core.SDKClient, req *legalcase.AlibabalegalcasemediaterecordsaveAPIRequest, session string) (*legalcase.AlibabalegalcasemediaterecordsaveAPIResponse, error) {
	var resp legalcase.AlibabalegalcasemediaterecordsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
