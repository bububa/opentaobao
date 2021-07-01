package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest
四级地址库-省 API请求
alibaba.onetouch.logistics.express.address.province.list

四级地址库-省 */
type AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// New
