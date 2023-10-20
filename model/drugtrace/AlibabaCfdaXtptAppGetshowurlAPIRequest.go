package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfdaxtptappgetshowurlAPIRequest 协同平台码查询页面url API请求
// alibaba.cfda.xtpt.app.getshowurl
//
// 协同平台码查询页面url
type AlibabacfdaxtptappgetshowurlAPIRequest struct {
	model.Params
	// 码
	_code string
}

// NewAlibabacfdaxtptappgetshowurlRequest 初始化AlibabacfdaxtptappgetshowurlAPIRequest对象
func NewAlibabacfdaxtptappgetshowurlRequest() *AlibabacfdaxtptappgetshowurlAPIRequest {
	return &AlibabacfdaxtptappgetshowurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacfdaxtptappgetshowurlAPIRequest) GetApiMethodName() string {
	return "alibaba.cfda.xtpt.app.getshowurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacfdaxtptappgetshowurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacfdaxtptappgetshowurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 码
func (r *AlibabacfdaxtptappgetshowurlAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabacfdaxtptappgetshowurlAPIRequest) GetCode() string {
	return r._code
}
