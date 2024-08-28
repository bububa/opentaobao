package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScUcrowdDelete 淘宝客-服务商-删除人群标签
// taobao.tbk.sc.ucrowd.delete
//
// 服务商使用。支持淘宝客删除人群标签id，被删除的人群标签id将失效，并不可重新生效。
func TaobaoTbkScUcrowdDelete(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScUcrowdDeleteAPIRequest, resp *tbk.TaobaoTbkScUcrowdDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
