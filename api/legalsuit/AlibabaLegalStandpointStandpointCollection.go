package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointCollection 收藏|取消收藏
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
func AlibabaLegalStandpointStandpointCollection(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointCollectionAPIRequest, resp *legalsuit.AlibabaLegalStandpointStandpointCollectionAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
