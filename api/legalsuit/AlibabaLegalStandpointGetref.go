package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointGetref 查询业务实体关联口径2
// alibaba.legal.standpoint.getref
//
// 口径查询
func AlibabaLegalStandpointGetref(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointGetrefAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointGetrefAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointGetrefAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
