package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGetrolebyempidAPIRequest
根据用户查询角色 API请求
alibaba.campus.acl.getrolebyempid

根据用户查询角色 */
type AlibabaCampusAclGetrolebyempidAPIRequest struct {
	model.Params
	// 公司id
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户id
	_param1 string
}

// New
