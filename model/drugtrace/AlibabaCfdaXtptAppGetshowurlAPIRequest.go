package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfdaXtptAppGetshowurlAPIRequest 协同平台码查询页面url API请求
// alibaba.cfda.xtpt.app.getshowurl
//
// 协同平台码查询页面url
type AlibabaCfdaXtptAppGetshowurlAPIRequest struct {
	model.Params
	// 码
	_code string
}

// NewAlibabaCfdaXtptAppGetshowurlRequest 初始化AlibabaCfdaXtptAppGetshowurlAPIRequest对象
func NewAlibabaCfdaXtptAppGetshowurlRequest() *AlibabaCfdaXtptAppGetshowurlAPIRequest {
	return &AlibabaCfdaXtptAppGetshowurlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetApiMethodName() string {
	return "alibaba.cfda.xtpt.app.getshowurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 码
func (r *AlibabaCfdaXtptAppGetshowurlAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetCode() string {
	return r._code
}
