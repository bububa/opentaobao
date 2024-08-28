package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// TaobaoXhotelComboReview 套餐审核接口
// taobao.xhotel.combo.review
//
// 套餐审核接口
func TaobaoXhotelComboReview(ctx context.Context, clt *core.SDKClient, req *tuanhotel.TaobaoXhotelComboReviewAPIRequest, resp *tuanhotel.TaobaoXhotelComboReviewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
