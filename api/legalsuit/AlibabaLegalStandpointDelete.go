package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointdelete 删除关联口径
// alibaba.legal.standpoint.delete
//
// 删除关联口径
func Alibabalegalstandpointdelete(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointdeleteAPIRequest, session string) (*legalsuit.AlibabalegalstandpointdeleteAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
