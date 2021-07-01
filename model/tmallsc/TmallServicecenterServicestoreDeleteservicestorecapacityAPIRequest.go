package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest
删除网点容量 API请求
tmall.servicecenter.servicestore.deleteservicestorecapacity

删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
必选字段：serviceStoreCode、bizType */
type TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest struct {
	model.Params
	// 网点编码
	_serviceStoreCode string
	// 业务类型
	_bizType string
}

// New
