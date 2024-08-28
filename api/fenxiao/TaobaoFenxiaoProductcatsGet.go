package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatsGet 查询产品线列表
// taobao.fenxiao.productcats.get
//
// 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
func TaobaoFenxiaoProductcatsGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoProductcatsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
