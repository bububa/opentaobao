package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointSceneQuery 查询场景
// alibaba.legal.standpoint.scene.query
//
// 查询场景
func AlibabaLegalStandpointSceneQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointSceneQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointSceneQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
