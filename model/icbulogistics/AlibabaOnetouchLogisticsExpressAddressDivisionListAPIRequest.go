package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest
四级地址库-区域 API请求
alibaba.onetouch.logistics.express.address.division.list

四级地址库-区 */
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// New
