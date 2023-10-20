package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstandpointgetref 查询业务实体关联口径2
// alibaba.legal.standpoint.getref
//
// 口径查询
func Alibabalegalstandpointgetref(clt *core.SDKClient, req *legalsuit.AlibabalegalstandpointgetrefAPIRequest, session string) (*legalsuit.AlibabalegalstandpointgetrefAPIResponse, error) {
	var resp legalsuit.AlibabalegalstandpointgetrefAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
