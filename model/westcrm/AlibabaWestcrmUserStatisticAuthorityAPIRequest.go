package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmUserStatisticAuthorityAPIRequest
获取指定用户是否含有会员权益统计权限 API请求
alibaba.westcrm.user.statistic.authority

获取指定用户是否含有会员权益统计权限 */
type AlibabaWestcrmUserStatisticAuthorityAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 用户id
	_ibUserId int64
}

// New
