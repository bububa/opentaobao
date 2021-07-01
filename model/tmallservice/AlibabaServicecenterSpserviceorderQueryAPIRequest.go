package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterSpserviceorderQueryAPIRequest
服务供应链服务单查询 API请求
alibaba.servicecenter.spserviceorder.query

服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务 */
type AlibabaServicecenterSpserviceorderQueryAPIRequest struct {
	model.Params
	// 状态码，可传多个
	_statusCodes string
	// 查询第几页，默认1
	_currentPage int64
	// 每页大小，默认50，最大50
	_pageSize int64
	// 服务单修改时间(时间段15分钟以内)(包含时分秒)
	_gmtModifiedEnd string
	// 服务单修改时间(包含时分秒)
	_gmtModifiedStart string
	// 实物主订单id(消费者在淘宝订单里看到的订单号)
	_masterParentBizOrderId int64
	// 服务单id
	_spServiceOrderId int64
}

// New
