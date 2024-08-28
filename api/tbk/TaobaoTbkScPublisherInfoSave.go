package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScPublisherInfoSave 淘宝客-公用-私域用户备案
// taobao.tbk.sc.publisher.info.save
//
// 通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
func TaobaoTbkScPublisherInfoSave(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScPublisherInfoSaveAPIRequest, resp *tbk.TaobaoTbkScPublisherInfoSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
