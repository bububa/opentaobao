package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemQuery 查询后端商品
// taobao.scitem.query
//
// 查询后端商品
func TaobaoScitemQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoScitemQueryAPIRequest, resp *fenxiao.TaobaoScitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
