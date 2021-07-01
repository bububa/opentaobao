package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest
四级地址库-市 API请求
alibaba.onetouch.logistics.express.address.city.list

四级地址库-市 */
type AlibabaOnetouchLogisticsExpressAddressCityListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// New
