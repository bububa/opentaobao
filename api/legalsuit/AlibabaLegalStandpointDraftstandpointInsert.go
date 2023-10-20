package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointdraftstandpointinsert 编辑后新增草稿口径
// alibaba.legal.standpoint.draftstandpoint.insert
//
// 编辑后新增草稿口径
func Alibabalegalstandpointdraftstandpointinsert(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointdraftstandpointinsertAPIRequest, session string) (*legalsuit.AlibabalegalstandpointdraftstandpointinsertAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointdraftstandpointinsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
