package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointCollection 收藏|取消收藏
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
func AlibabaLegalStandpointStandpointCollection(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointCollectionAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointStandpointCollectionAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointStandpointCollectionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
