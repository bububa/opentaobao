package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointDelete 删除关联口径
// alibaba.legal.standpoint.delete
//
// 删除关联口径
func AlibabaLegalStandpointDelete(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointDeleteAPIRequest, resp *legalsuit.AlibabaLegalStandpointDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
