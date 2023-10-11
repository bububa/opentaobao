package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointQueryall 滑动查询口径
// alibaba.legal.standpoint.standpoint.queryall
//
// 滑动查询口径
func AlibabaLegalStandpointStandpointQueryall(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointQueryallAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointStandpointQueryallAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointStandpointQueryallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
