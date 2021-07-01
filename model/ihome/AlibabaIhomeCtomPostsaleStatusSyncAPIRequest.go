package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomPostsaleStatusSyncAPIRequest
C2M售后状态同步 API请求
alibaba.ihome.ctom.postsale.status.sync

供给三维家同步定制、成品商品售后进度状态 */
type AlibabaIhomeCtomPostsaleStatusSyncAPIRequest struct {
	model.Params
	// 三维家服务ID
	_serviceId string
	// 三维家售后单号ID
	_postSalesId string
	// 三维家订单号ID
	_subOrderId string
	// 三维家操作人部门ID
	_unitId string
	// 三维家操作人ID
	_operatorId string
	// 售后状态更新
	_status string
	// 售后发起来源
	_source string
	// 是否加急订单
	_isExpedited string
	// 售后单更新状态原因
	_reason string
	// 售后单结束原因
	_finishType string
	// 客服代表ID
	_csrId string
}

// New
