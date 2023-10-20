package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointInsertdraft 插入草稿
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
func AlibabaLegalStandpointInsertdraft(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointInsertdraftAPIRequest, resp *legalsuit.AlibabaLegalStandpointInsertdraftAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
