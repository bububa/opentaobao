package dengta

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaWbaccountPriceChange 微博公众号价格变化通知
// alibaba.pictures.dengta.wbaccount.price.change
//
// 微博公众号推广价格变更通知接口
func AlibabaPicturesDengtaWbaccountPriceChange(ctx context.Context, clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaWbaccountPriceChangeAPIRequest, resp *dengta.AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
