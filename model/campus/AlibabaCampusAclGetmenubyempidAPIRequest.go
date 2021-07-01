package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGetmenubyempidAPIRequest
查询用户的菜单 API请求
alibaba.campus.acl.getmenubyempid

查询用户的菜单 */
type AlibabaCampusAclGetmenubyempidAPIRequest struct {
	model.Params
	// 账户id
	_userId int64
	// 系统id
	_systemId string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// New
