package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentcompanyinfogetAPIRequest 获取厂牌信息 API请求
// xiami.content.company.info.get
//
// 获取厂牌信息
type XiamicontentcompanyinfogetAPIRequest struct {
	model.Params
	// 厂牌id
	_companyIds int64
}

// NewXiamicontentcompanyinfogetRequest 初始化XiamicontentcompanyinfogetAPIRequest对象
func NewXiamicontentcompanyinfogetRequest() *XiamicontentcompanyinfogetAPIRequest {
	return &XiamicontentcompanyinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentcompanyinfogetAPIRequest) GetApiMethodName() string {
	return "xiami.content.company.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentcompanyinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentcompanyinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyIds is CompanyIds Setter
// 厂牌id
func (r *XiamicontentcompanyinfogetAPIRequest) SetCompanyIds(_companyIds int64) error {
	r._companyIds = _companyIds
	r.Set("company_ids", _companyIds)
	return nil
}

// GetCompanyIds CompanyIds Getter
func (r XiamicontentcompanyinfogetAPIRequest) GetCompanyIds() int64 {
	return r._companyIds
}
