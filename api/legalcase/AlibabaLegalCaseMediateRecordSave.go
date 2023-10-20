package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseMediateRecordSave 新增调解结果
// alibaba.legal.case.mediate.record.save
//
// 增加调解沟通记录
func AlibabaLegalCaseMediateRecordSave(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseMediateRecordSaveAPIRequest, resp *legalcase.AlibabaLegalCaseMediateRecordSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
