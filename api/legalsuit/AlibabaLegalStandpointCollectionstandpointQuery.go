package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointCollectionstandpointQuery 查询收藏口径
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
func AlibabaLegalStandpointCollectionstandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointCollectionstandpointQueryAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointCollectionstandpointQueryAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointCollectionstandpointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
