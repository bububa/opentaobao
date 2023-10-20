package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusvisitorimageuploadAPIRequest 访客大厅图片上传及查看 API请求
// alibaba.campus.visitor.image.upload
//
// 访客大厅图片上传及查看
type AlibabacampusvisitorimageuploadAPIRequest struct {
	model.Params
	// 无意义参数
	_noneString string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabacampusvisitorimageuploadRequest 初始化AlibabacampusvisitorimageuploadAPIRequest对象
func NewAlibabacampusvisitorimageuploadRequest() *AlibabacampusvisitorimageuploadAPIRequest {
	return &AlibabacampusvisitorimageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusvisitorimageuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.visitor.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusvisitorimageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusvisitorimageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNoneString is NoneString Setter
// 无意义参数
func (r *AlibabacampusvisitorimageuploadAPIRequest) SetNoneString(_noneString string) error {
	r._noneString = _noneString
	r.Set("none_string", _noneString)
	return nil
}

// GetNoneString NoneString Getter
func (r AlibabacampusvisitorimageuploadAPIRequest) GetNoneString() string {
	return r._noneString
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampusvisitorimageuploadAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusvisitorimageuploadAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusvisitorimageuploadAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusvisitorimageuploadAPIRequest) GetCampusId() int64 {
	return r._campusId
}
