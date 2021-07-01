package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest
权益池信息查询 API请求
alibaba.interact.supplier.award.resource.get.cuntao

农村淘宝营销互动权益池信息查询 */
type AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest struct {
	model.Params
	// 用户昵称
	_userNick string
	// 活动code
	_activityKey string
	// 经度
	_lng string
	// 纬度
	_lat string
}

// New
