package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScPublisherInfoGet 淘宝客-公用-私域用户备案信息查询
// taobao.tbk.sc.publisher.info.get
//
// 查询已生成的渠道id或会员运营id的相关信息。
func TaobaoTbkScPublisherInfoGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScPublisherInfoGetAPIRequest, resp *tbk.TaobaoTbkScPublisherInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
