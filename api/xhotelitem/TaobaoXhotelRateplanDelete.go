package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanDelete 价格计划rateplan删除
// taobao.xhotel.rateplan.delete
//
// 酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
func TaobaoXhotelRateplanDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelRateplanDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
