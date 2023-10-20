package tmallcampus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcampusauthstatusqueryAPIRequest 学生认证状态查询 API请求
// tmall.campus.authstatus.query
//
// 学生认证状态查询
type TmallcampusauthstatusqueryAPIRequest struct {
	model.Params
}

// NewTmallcampusauthstatusqueryRequest 初始化TmallcampusauthstatusqueryAPIRequest对象
func NewTmallcampusauthstatusqueryRequest() *TmallcampusauthstatusqueryAPIRequest {
	return &TmallcampusauthstatusqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcampusauthstatusqueryAPIRequest) GetApiMethodName() string {
	return "tmall.campus.authstatus.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcampusauthstatusqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcampusauthstatusqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
