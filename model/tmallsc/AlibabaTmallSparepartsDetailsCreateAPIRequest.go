package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallsparepartsdetailscreateAPIRequest 天猫蚁巢同步工单申请备件明细 API请求
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
type AlibabatmallsparepartsdetailscreateAPIRequest struct {
	model.Params
	// 请求入参
	_sparePartsDetailsSaveRequest *SparePartsDetailsSaveRequest
}

// NewAlibabatmallsparepartsdetailscreateRequest 初始化AlibabatmallsparepartsdetailscreateAPIRequest对象
func NewAlibabatmallsparepartsdetailscreateRequest() *AlibabatmallsparepartsdetailscreateAPIRequest {
	return &AlibabatmallsparepartsdetailscreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallsparepartsdetailscreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tmall.spareparts.details.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallsparepartsdetailscreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallsparepartsdetailscreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSparePartsDetailsSaveRequest is SparePartsDetailsSaveRequest Setter
// 请求入参
func (r *AlibabatmallsparepartsdetailscreateAPIRequest) SetSparePartsDetailsSaveRequest(_sparePartsDetailsSaveRequest *SparePartsDetailsSaveRequest) error {
	r._sparePartsDetailsSaveRequest = _sparePartsDetailsSaveRequest
	r.Set("spare_parts_details_save_request", _sparePartsDetailsSaveRequest)
	return nil
}

// GetSparePartsDetailsSaveRequest SparePartsDetailsSaveRequest Getter
func (r AlibabatmallsparepartsdetailscreateAPIRequest) GetSparePartsDetailsSaveRequest() *SparePartsDetailsSaveRequest {
	return r._sparePartsDetailsSaveRequest
}
