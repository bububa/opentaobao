package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpToolGet 查询工具
// taobao.ump.tool.get
//
// 根据工具id获取一个工具对象
func TaobaoUmpToolGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpToolGetAPIRequest, resp *promotion.TaobaoUmpToolGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
