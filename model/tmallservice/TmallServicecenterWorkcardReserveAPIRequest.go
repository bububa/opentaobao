package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardReserveAPIRequest
工单预约 API请求
tmall.servicecenter.workcard.reserve

服务工单更新通用接口 */
type TmallServicecenterWorkcardReserveAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 服务开始时间
	_reserveTimeStart string
	// 服务结束时间
	_reserveTimeEnd string
	// 工人手机号
	_workerMobile string
	// 预约备注信息
	_reserveRemark string
	// 工人姓名
	_workerName string
	// 扩展信息
	_attributes string
	// 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
	_action string
	// 履约类型，上门或者到店
	_fulfilType string
	// 门店编码
	_serviceStoreCode string
	// 门店名称
	_serviceStoreName string
	// 核销单id
	_fulfilTaskId int64
}

// New
