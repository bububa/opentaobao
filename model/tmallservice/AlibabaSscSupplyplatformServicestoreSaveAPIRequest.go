package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServicestoreSaveAPIRequest
保存网点 API请求
alibaba.ssc.supplyplatform.servicestore.save

网点创建、修改 */
type AlibabaSscSupplyplatformServicestoreSaveAPIRequest struct {
	model.Params
	// 入参
	_serviceStoreSaveReq *ServiceStoreSaveForTopReqDto
}

// New
