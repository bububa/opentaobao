package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreCreateAPIRequest
创建门店 API请求
tmall.servicecenter.servicestore.create

用于创建门店/网点。多个业务共用 */
type TmallServicecenterServicestoreCreateAPIRequest struct {
	model.Params
	// 网点/门店
	_serviceStore *ServiceStoreDto
}

// New
