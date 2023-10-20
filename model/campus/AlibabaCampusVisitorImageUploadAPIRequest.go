package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusVisitorImageUploadAPIRequest 访客大厅图片上传及查看 API请求
// alibaba.campus.visitor.image.upload
//
// 访客大厅图片上传及查看
type AlibabaCampusVisitorImageUploadAPIRequest struct {
	model.Params
	// 无意义参数
	_noneString string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabaCampusVisitorImageUploadRequest 初始化AlibabaCampusVisitorImageUploadAPIRequest对象
func NewAlibabaCampusVisitorImageUploadRequest() *AlibabaCampusVisitorImageUploadAPIRequest {
	return &AlibabaCampusVisitorImageUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusVisitorImageUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.visitor.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusVisitorImageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusVisitorImageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNoneString is NoneString Setter
// 无意义参数
func (r *AlibabaCampusVisitorImageUploadAPIRequest) SetNoneString(_noneString string) error {
	r._noneString = _noneString
	r.Set("none_string", _noneString)
	return nil
}

// GetNoneString NoneString Getter
func (r AlibabaCampusVisitorImageUploadAPIRequest) GetNoneString() string {
	return r._noneString
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabaCampusVisitorImageUploadAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusVisitorImageUploadAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusVisitorImageUploadAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusVisitorImageUploadAPIRequest) GetCampusId() int64 {
	return r._campusId
}
