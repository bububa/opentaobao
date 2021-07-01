package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderNoticeAPIRequest
手动触发发短信 API请求
cainiao.endpoint.locker.top.order.notice

合作公司对订单手动触发短信，有次数限制 */
type CainiaoEndpointLockerTopOrderNoticeAPIRequest struct {
	model.Params
	// 合作公司唯一订单编号
	_orderCode string
	// 站点ID
	_stationId string
	// 运单号
	_mailNo string
	// 场景编号：0：重发短信，1：催取短信
	_sceneCode int64
}

// New
