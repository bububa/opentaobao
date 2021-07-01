package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmUserConsumerGetAPIRequest
获取指定用户的消费总额 API请求
alibaba.westcrm.user.consumer.get

获取指定用户的消费总额 */
type AlibabaWestcrmUserConsumerGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 用户id
	_ibUserId int64
	// 开始时间
	_timeBegin string
	// 结束时间
	_timeEnd string
}

// New
