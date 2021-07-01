package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmUserBenefitAuthorityAPIRequest
获取指定用户是否含有会员权益配置菜单权限 API请求
alibaba.westcrm.user.benefit.authority

获取指定用户是否含有会员权益配置菜单权限 */
type AlibabaWestcrmUserBenefitAuthorityAPIRequest struct {
	model.Params
	// 园区ID
	_campusId int64
	// 当前用户id
	_ibUserId int64
}

// New
