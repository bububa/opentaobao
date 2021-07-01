package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreUpdateAPIRequest
修改门店信息 API请求
tmall.servicecenter.servicestore.update

用于修改门店/网点信息。多个业务共用 */
type TmallServicecenterServicestoreUpdateAPIRequest struct {
	model.Params
	// 网点/门店
	_paramServiceStoreDTO *ServiceStoreDto
}

// New
