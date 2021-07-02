package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest 同城零售履约异常中心异常单处理结果回调接口 API请求
// tmall.cityretail.fulfill.abnormal.center.abnormal.status.change
//
// 同城零售履约异常中心异常单处理结果回调接口
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest struct {
	model.Params
	// 入参
	_abnormalStatusChangeDto []AbnormalStatusChangeDto
}

// NewTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest 初始化TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest对象
func NewTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest() *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest {
	return &TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.fulfill.abnormal.center.abnormal.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AbnormalStatusChangeDto Setter
// 入参
func (r *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest) SetAbnormalStatusChangeDto(_abnormalStatusChangeDto []AbnormalStatusChangeDto) error {
	r._abnormalStatusChangeDto = _abnormalStatusChangeDto
	r.Set("abnormal_status_change_dto", _abnormalStatusChangeDto)
	return nil
}

// Get AbnormalStatusChangeDto Getter
func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest) GetAbnormalStatusChangeDto() []AbnormalStatusChangeDto {
	return r._abnormalStatusChangeDto
}
