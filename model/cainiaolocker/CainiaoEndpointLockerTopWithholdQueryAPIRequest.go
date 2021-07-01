package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopWithholdQueryAPIRequest
查询能否代扣 API请求
cainiao.endpoint.locker.top.withhold.query

查询是否有代扣欠款，是否签署代扣协议。 */
type CainiaoEndpointLockerTopWithholdQueryAPIRequest struct {
	model.Params
	// 柜子公司编码
	_companyCode string
	// 开放用户Id
	_openUserId string
	// 柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用）
	_orderType int64
}

// New
