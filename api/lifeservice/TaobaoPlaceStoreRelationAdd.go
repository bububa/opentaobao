package lifeservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lifeservice"
)

// TaobaoPlaceStoreRelationAdd 门店关系新增
// taobao.place.store.relation.add
//
// 新增授权用户的门店关系信息
func TaobaoPlaceStoreRelationAdd(ctx context.Context, clt *core.SDKClient, req *lifeservice.TaobaoPlaceStoreRelationAddAPIRequest, resp *lifeservice.TaobaoPlaceStoreRelationAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
