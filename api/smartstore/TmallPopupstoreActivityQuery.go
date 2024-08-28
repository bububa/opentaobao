package smartstore

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TmallPopupstoreActivityQuery 查询某段时间内的快闪活动列表
// tmall.popupstore.activity.query
//
// 提供给ISV查询某一时间段内包含指定appKey的活动列表
func TmallPopupstoreActivityQuery(ctx context.Context, clt *core.SDKClient, req *smartstore.TmallPopupstoreActivityQueryAPIRequest, resp *smartstore.TmallPopupstoreActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
