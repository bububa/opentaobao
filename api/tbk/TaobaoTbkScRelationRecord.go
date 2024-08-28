package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScRelationRecord 淘宝客-服务商-私域会员标记
// taobao.tbk.sc.relation.record
//
// 服务商使用。支持淘宝客通过入参私域外部ID，获得待私域会员可标记的链接，会员打开该链接后，可帮助媒体自动生成会员运营id进行标记，同时自动跳转到推广落地页。
func TaobaoTbkScRelationRecord(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScRelationRecordAPIRequest, resp *tbk.TaobaoTbkScRelationRecordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
