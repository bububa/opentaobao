package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressDivisionList 四级地址库-区域
// alibaba.onetouch.logistics.express.address.division.list
//
// 四级地址库-区
func AlibabaOnetouchLogisticsExpressAddressDivisionList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
