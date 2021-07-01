package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServicestoreOfflineAPIRequest
网点下线 API请求
alibaba.ssc.supplyplatform.servicestore.offline

网点下线功能 */
type AlibabaSscSupplyplatformServicestoreOfflineAPIRequest struct {
	model.Params
	// 网点编码列表集合,最大支持1000
	_serviceStoreCodeList []string
	// 网点id列表集合,最大支持1000
	_serviceStoreIdList []int64
}

// New
