package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCfdaXtptAppGetshowurlAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetApiMethodName() string {
	return "alibaba.cfda.xtpt.app.getshowurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 码
func (r *AlibabaCfdaXtptAppGetshowurlAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaCfdaXtptAppGetshowurlAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaCfdaXtptAppGetshowurlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCfdaXtptAppGetshowurlRequest()
	},
}

// GetAlibabaCfdaXtptAppGetshowurlRequest 从 sync.Pool 获取 AlibabaCfdaXtptAppGetshowurlAPIRequest
func GetAlibabaCfdaXtptAppGetshowurlAPIRequest() *AlibabaCfdaXtptAppGetshowurlAPIRequest {
	return poolAlibabaCfdaXtptAppGetshowurlAPIRequest.Get().(*AlibabaCfdaXtptAppGetshowurlAPIRequest)
}

// ReleaseAlibabaCfdaXtptAppGetshowurlAPIRequest 将 AlibabaCfdaXtptAppGetshowurlAPIRequest 放入 sync.Pool
func ReleaseAlibabaCfdaXtptAppGetshowurlAPIRequest(v *AlibabaCfdaXtptAppGetshowurlAPIRequest) {
	v.Reset()
	poolAlibabaCfdaXtptAppGetshowurlAPIRequest.Put(v)
}
