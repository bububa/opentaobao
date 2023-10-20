package xiamicontent

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentCompanyInfoGetAPIRequest 获取厂牌信息 API请求
// xiami.content.company.info.get
//
// 获取厂牌信息
type XiamiContentCompanyInfoGetAPIRequest struct {
	model.Params
	// 厂牌id
	_companyIds int64
}

// NewXiamiContentCompanyInfoGetRequest 初始化XiamiContentCompanyInfoGetAPIRequest对象
func NewXiamiContentCompanyInfoGetRequest() *XiamiContentCompanyInfoGetAPIRequest {
	return &XiamiContentCompanyInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiContentCompanyInfoGetAPIRequest) Reset() {
	r._companyIds = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentCompanyInfoGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.company.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentCompanyInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentCompanyInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyIds is CompanyIds Setter
// 厂牌id
func (r *XiamiContentCompanyInfoGetAPIRequest) SetCompanyIds(_companyIds int64) error {
	r._companyIds = _companyIds
	r.Set("company_ids", _companyIds)
	return nil
}

// GetCompanyIds CompanyIds Getter
func (r XiamiContentCompanyInfoGetAPIRequest) GetCompanyIds() int64 {
	return r._companyIds
}

var poolXiamiContentCompanyInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiContentCompanyInfoGetRequest()
	},
}

// GetXiamiContentCompanyInfoGetRequest 从 sync.Pool 获取 XiamiContentCompanyInfoGetAPIRequest
func GetXiamiContentCompanyInfoGetAPIRequest() *XiamiContentCompanyInfoGetAPIRequest {
	return poolXiamiContentCompanyInfoGetAPIRequest.Get().(*XiamiContentCompanyInfoGetAPIRequest)
}

// ReleaseXiamiContentCompanyInfoGetAPIRequest 将 XiamiContentCompanyInfoGetAPIRequest 放入 sync.Pool
func ReleaseXiamiContentCompanyInfoGetAPIRequest(v *XiamiContentCompanyInfoGetAPIRequest) {
	v.Reset()
	poolXiamiContentCompanyInfoGetAPIRequest.Put(v)
}
