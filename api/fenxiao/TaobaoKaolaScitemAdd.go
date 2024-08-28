package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoKaolaScitemAdd 考拉货品新增接口
// taobao.kaola.scitem.add
//
// 考拉货品新增接口
func TaobaoKaolaScitemAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoKaolaScitemAddAPIRequest, resp *fenxiao.TaobaoKaolaScitemAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
