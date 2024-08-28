package icbulogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressStreetList 四级地址库-街道
// alibaba.onetouch.logistics.express.address.street.list
//
// 四级地址库-街道模糊查询
func AlibabaOnetouchLogisticsExpressAddressStreetList(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressStreetListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
