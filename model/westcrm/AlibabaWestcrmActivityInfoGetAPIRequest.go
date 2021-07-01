package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmActivityInfoGetAPIRequest
获取活动详情 API请求
alibaba.westcrm.activity.info.get

根据id查询活动详情 */
type AlibabaWestcrmActivityInfoGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 活动id
	_activityId int64
}

// New
