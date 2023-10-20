package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointinsertdraft 插入草稿
// alibaba.legal.standpoint.insertdraft
//
// 插入草稿
func Alibabalegalstandpointinsertdraft(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointinsertdraftAPIRequest, session string) (*legalsuit.AlibabalegalstandpointinsertdraftAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointinsertdraftAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
