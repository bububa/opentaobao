package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTaxationCalculateOpenQueryAPIRequest
关务所需的申报清关字段 API请求
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段 */
type AliexpressTaxationCalculateOpenQueryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId string
}

// New
