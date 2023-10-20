package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressProvinceList 四级地址库-省
// alibaba.onetouch.logistics.express.address.province.list
//
// 四级地址库-省
func AlibabaOnetouchLogisticsExpressAddressProvinceList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
