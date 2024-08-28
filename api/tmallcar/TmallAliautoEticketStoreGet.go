package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoEticketStoreGet 查询电子凭证对应门店信息
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
func TmallAliautoEticketStoreGet(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoEticketStoreGetAPIRequest, resp *tmallcar.TmallAliautoEticketStoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
