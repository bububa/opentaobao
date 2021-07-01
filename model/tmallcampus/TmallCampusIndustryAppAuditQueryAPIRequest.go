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

// NewTmallCampusIndustryAppAuditQueryRequest 初始化TmallCampusIndustryAppAuditQueryAPIRequest对象
func NewTmallCampusIndustryAppAuditQueryRequest() *TmallCampusIndustryAppAuditQueryAPIRequest {
	return &TmallCampusIndustryAppAuditQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCampusIndustryAppAuditQueryAPIRequest) GetApiMethodName() string {
	return "tmall.campus.industry.app.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCampusIndustryAppAuditQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Source Setter
// 调用来源
func (r *TmallCampusIndustryAppAuditQueryAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r TmallCampusIndustryAppAuditQueryAPIRequest) GetSource() string {
	return r._source
}
