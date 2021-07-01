package tmallcampus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCampusIndustryAppAuditQueryAPIRequest
天猫校园查询学生认证信息 API请求
tmall.campus.industry.app.audit.query

天猫校园查询学生认证信息 */
type TmallCampusIndustryAppAuditQueryAPIRequest struct {
	model.Params
	// 调用来源
	_source string
}

// New
