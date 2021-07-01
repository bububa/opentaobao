package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalStoresanditemsAPIRequest
查询商户门店，商品列表 API请求
alibaba.alihealth.reserve.dental.storesanditems

查询商户门店，商品列表 */
type AlibabaAlihealthReserveDentalStoresanditemsAPIRequest struct {
	model.Params
	// 页码，每页100个门店，超过100个门店分页请求
	_pageNo int64
}

// New
