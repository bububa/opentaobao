package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductAdd 添加产品
// taobao.fenxiao.product.add
//
// 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。&lt;br/&gt;&lt;br/&gt;    * 产品图片默认为空&lt;br/&gt;    * 产品发布后默认为下架状态
func TaobaoFenxiaoProductAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductAddAPIRequest, resp *fenxiao.TaobaoFenxiaoProductAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
