package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest
查询订单是否由裹裹发送消息 API请求
cainiao.endpoint.locker.top.order.noticesend.query

合作公司查询消息发送的接口，判断是否裹裹发送消息 */
type CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest struct {
	model.Params
	// 站点id
	_stationId string
	// 收件人手机号
	_getterPhone string
	// 运单号
	_mailNo string
}

// New
