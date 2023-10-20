package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointGetref 查询业务实体关联口径2
// alibaba.legal.standpoint.getref
//
// 口径查询
func AlibabaLegalStandpointGetref(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointGetrefAPIRequest, resp *legalsuit.AlibabaLegalStandpointGetrefAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
