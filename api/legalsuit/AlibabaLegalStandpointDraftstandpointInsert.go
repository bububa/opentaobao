package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointDraftstandpointInsert 编辑后新增草稿口径
// alibaba.legal.standpoint.draftstandpoint.insert
//
// 编辑后新增草稿口径
func AlibabaLegalStandpointDraftstandpointInsert(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointDraftstandpointInsertAPIRequest, resp *legalsuit.AlibabaLegalStandpointDraftstandpointInsertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
