package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderGuideDataGet 获取全渠道导购产品数据
// taobao.omniorder.guide.data.get
//
// 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
func TaobaoOmniorderGuideDataGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderGuideDataGetAPIRequest, resp *omniorder.TaobaoOmniorderGuideDataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
