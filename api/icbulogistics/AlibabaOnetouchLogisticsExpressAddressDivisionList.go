package icbulogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressDivisionList 四级地址库-区域
// alibaba.onetouch.logistics.express.address.division.list
//
// 四级地址库-区
func AlibabaOnetouchLogisticsExpressAddressDivisionList(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
