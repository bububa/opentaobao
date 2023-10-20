package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointDelete 删除关联口径
// alibaba.legal.standpoint.delete
//
// 删除关联口径
func AlibabaLegalStandpointDelete(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointDeleteAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointDeleteAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
