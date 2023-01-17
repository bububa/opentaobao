package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointInsertdraft 插入草稿
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
func AlibabaLegalStandpointInsertdraft(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointInsertdraftAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointInsertdraftAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointInsertdraftAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
