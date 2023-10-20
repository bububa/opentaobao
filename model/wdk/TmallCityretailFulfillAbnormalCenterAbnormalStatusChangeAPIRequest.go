package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest 同城零售履约异常中心异常单处理结果回调接口 API请求
// tmall.cityretail.fulfill.abnormal.center.abnormal.status.change
//
// 同城零售履约异常中心异常单处理结果回调接口
type TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest struct {
	model.Params
	// 入参
	_abnormalStatusChangeDto []AbnormalStatusChangeDto
}

// NewTmallcityretailfulfillabnormalcenterabnormalstatuschangeRequest 初始化TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest对象
func NewTmallcityretailfulfillabnormalcenterabnormalstatuschangeRequest() *TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest {
	return &TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.fulfill.abnormal.center.abnormal.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAbnormalStatusChangeDto is AbnormalStatusChangeDto Setter
// 入参
func (r *TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest) SetAbnormalStatusChangeDto(_abnormalStatusChangeDto []AbnormalStatusChangeDto) error {
	r._abnormalStatusChangeDto = _abnormalStatusChangeDto
	r.Set("abnormal_status_change_dto", _abnormalStatusChangeDto)
	return nil
}

// GetAbnormalStatusChangeDto AbnormalStatusChangeDto Getter
func (r TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest) GetAbnormalStatusChangeDto() []AbnormalStatusChangeDto {
	return r._abnormalStatusChangeDto
}
