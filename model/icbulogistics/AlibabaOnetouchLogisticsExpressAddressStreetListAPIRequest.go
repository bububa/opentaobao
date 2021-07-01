package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest
四级地址库-街道 API请求
alibaba.onetouch.logistics.express.address.street.list

四级地址库-街道模糊查询 */
type AlibabaOnetouchLogisticsExpressAddressStreetListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// New
