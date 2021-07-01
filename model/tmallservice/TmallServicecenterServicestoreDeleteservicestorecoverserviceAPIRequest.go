package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest
删除网点覆盖的服务 API请求
tmall.servicecenter.servicestore.deleteservicestorecoverservice

天猫服务平台删除网点覆盖的服务，
必选字段：serviceStoreCode、bizType */
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest struct {
	model.Params
	// 网点编码
	_serviceStoreCode string
	// 业务类型
	_bizType string
}

// New
