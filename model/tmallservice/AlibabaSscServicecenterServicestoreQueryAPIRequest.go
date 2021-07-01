package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscServicecenterServicestoreQueryAPIRequest
根据天猫id查询门店信息 API请求
alibaba.ssc.servicecenter.servicestore.query

根据天猫id查询门店信息 */
type AlibabaSscServicecenterServicestoreQueryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
}

// New
