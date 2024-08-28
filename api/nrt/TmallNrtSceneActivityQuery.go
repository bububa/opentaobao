package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtSceneActivityQuery 喵零场景活动查询
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
func TmallNrtSceneActivityQuery(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtSceneActivityQueryAPIRequest, resp *nrt.TmallNrtSceneActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
