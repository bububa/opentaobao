package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCharityCharitytimeQueryAPIRequest
查询公益3小时公益时汇总 API请求
alibaba.charity.charitytime.query

查询公益3小时公益时汇总 */
type AlibabaCharityCharitytimeQueryAPIRequest struct {
	model.Params
	// 公益类型
	_charityTypeList []string
	// 结束时间戳-毫秒时间
	_endDate int64
	// 开始时间戳-毫秒时间
	_startDate int64
	// 淘宝Uid
	_tbUid int64
	// 活动ID
	_activityId int64
	// 扩展参数
	_extParam string
}

// New
