package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemstoreBanding 商品关联绑定接口
// taobao.qimen.itemstore.banding
//
// 商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
func TaobaoQimenItemstoreBanding(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenItemstoreBandingAPIRequest, resp *qimen.TaobaoQimenItemstoreBandingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
