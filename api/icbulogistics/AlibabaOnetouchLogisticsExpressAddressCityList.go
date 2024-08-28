package icbulogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressCityList 四级地址库-市
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
func AlibabaOnetouchLogisticsExpressAddressCityList(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
