package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMallitemcenterSupplierPriceUploadAPIRequest
天猫服务商服务报价上传 API请求
tmall.mallitemcenter.supplier.price.upload

天猫服务商上传服务价格 */
type TmallMallitemcenterSupplierPriceUploadAPIRequest struct {
	model.Params
	// 服务code
	_serviceCode string
	// 服务商门店价格列表
	_providerPriceList []StoreOfferPriceDto
}

// New
