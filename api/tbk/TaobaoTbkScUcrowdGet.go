package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScUcrowdGet 淘宝客-服务商-获取人群标签
// taobao.tbk.sc.ucrowd.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群信息，包括人群名称描述及覆盖会员数。
func TaobaoTbkScUcrowdGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScUcrowdGetAPIRequest, resp *tbk.TaobaoTbkScUcrowdGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
