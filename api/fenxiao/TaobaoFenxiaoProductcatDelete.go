package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatDelete 删除产品线
// taobao.fenxiao.productcat.delete
//
// 删除产品线
func TaobaoFenxiaoProductcatDelete(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatDeleteAPIRequest, resp *fenxiao.TaobaoFenxiaoProductcatDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
