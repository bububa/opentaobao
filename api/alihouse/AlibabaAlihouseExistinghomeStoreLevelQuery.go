package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreLevelQuery 门店等级评分查询
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
func AlibabaAlihouseExistinghomeStoreLevelQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
