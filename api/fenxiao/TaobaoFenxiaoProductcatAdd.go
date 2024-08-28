package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatAdd 新增产品线
// taobao.fenxiao.productcat.add
//
// 新增产品线
func TaobaoFenxiaoProductcatAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatAddAPIRequest, resp *fenxiao.TaobaoFenxiaoProductcatAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
