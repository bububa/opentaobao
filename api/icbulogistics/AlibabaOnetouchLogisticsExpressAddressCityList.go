package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressCityList 四级地址库-市
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
func AlibabaOnetouchLogisticsExpressAddressCityList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
