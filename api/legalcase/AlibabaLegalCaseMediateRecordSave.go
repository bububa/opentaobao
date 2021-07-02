package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseMediateRecordSave 新增调解结果
// alibaba.legal.case.mediate.record.save
//
// 增加调解沟通记录
func AlibabaLegalCaseMediateRecordSave(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseMediateRecordSaveAPIRequest, session string) (*legalcase.AlibabaLegalCaseMediateRecordSaveAPIResponse, error) {
	var resp legalcase.AlibabaLegalCaseMediateRecordSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
