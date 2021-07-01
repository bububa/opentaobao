package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalItemBindAPIRequest
ISV绑定外部门店id和外部商品id API请求
alibaba.alihealth.dental.item.bind

ISV绑定外部门店id和外部商品id */
type AlibabaAlihealthDentalItemBindAPIRequest struct {
	model.Params
	// bind_list
	_bindList []StoreItemRelRequest
	// 类型 1 天猫门店 2 支付宝门店
	_type int64
}

// New
