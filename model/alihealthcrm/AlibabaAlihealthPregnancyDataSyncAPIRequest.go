package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyDataSyncAPIRequest
四类数据同步 API请求
alibaba.alihealth.pregnancy.data.sync

经期调整；基础体温；排卵试纸；B超测排数据同步 */
type AlibabaAlihealthPregnancyDataSyncAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
	// 三方id
	_outerId int64
	// 4经期调整；1基础体温；2排卵试纸；3 B超测排
	_eventType int64
	// 四类数据定制化详情
	_data string
	// 测量日期
	_measureDate int64
	// 0-新增 1-修改 2-删除
	_operationType int64
	// 经期数据json串
	_periodMsg string
}

// New
