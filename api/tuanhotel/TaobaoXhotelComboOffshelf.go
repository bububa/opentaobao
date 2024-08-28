package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// TaobaoXhotelComboOffshelf 酒店套餐下架
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
func TaobaoXhotelComboOffshelf(ctx context.Context, clt *core.SDKClient, req *tuanhotel.TaobaoXhotelComboOffshelfAPIRequest, resp *tuanhotel.TaobaoXhotelComboOffshelfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
