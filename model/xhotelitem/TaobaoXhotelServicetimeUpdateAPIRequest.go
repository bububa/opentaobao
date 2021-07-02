package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelServicetimeUpdateAPIRequest 飞猪酒店多维度服务时间维护接口 API请求
// taobao.xhotel.servicetime.update
//
// 飞猪酒店多维度服务时间维护，支持卖家维度，supplier维度，酒店维度
type TaobaoXhotelServicetimeUpdateAPIRequest struct {
	model.Params
	// 请按照示例值的格式来填写，涉及到是否当日订单，是否展示，周一到周日的服务时间，业务id,业务类型1为卖家，2为supplier ，3为酒店。[{"businessId":11925099374,"businessType":3,"displayItemInNonWorkingTime":1,"mondayConfirmLocalTime":"09:00-18:00","operator":"操作人","orderConfirmType":1,"saturdayConfirmLocalTime":"09:00-18:00","sellerId":2054718374,"sellerNick":"sandbox_b_27","sundayConfirmLocalTime":"09:00-18:00","supplier":"","thursdayConfirmLocalTime":"09:00-18:00","timeZoneName":"Asia/Shanghai","tuesdayConfirmLocalTime":"09:00-18:00","wednesdayConfirmLocalTime":"09:00-18:00","fridayConfirmLocalTime":"09:00-18:00"},{"businessId":11925099374,"businessType":3,"displayItemInNonWorkingTime":1,"mondayConfirmLocalTime":"09:00-18:00","operator":"操作人","orderConfirmType":2,"saturdayConfirmLocalTime":"09:00-18:00","sellerId":2054718374,"sellerNick":"sandbox_b_27","sundayConfirmLocalTime":"09:00-18:00","supplier":"","thursdayConfirmLocalTime":"09:00-18:00","timeZoneName":"Asia/Shanghai","tuesdayConfirmLocalTime":"09:00-18:00","wednesdayConfirmLocalTime":"09:00-18:00","fridayConfirmLocalTime":"09:00-18:00"}]
	_param string
}

// NewTaobaoXhotelServicetimeUpdateRequest 初始化TaobaoXhotelServicetimeUpdateAPIRequest对象
func NewTaobaoXhotelServicetimeUpdateRequest() *TaobaoXhotelServicetimeUpdateAPIRequest {
	return &TaobaoXhotelServicetimeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.servicetime.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 请按照示例值的格式来填写，涉及到是否当日订单，是否展示，周一到周日的服务时间，业务id,业务类型1为卖家，2为supplier ，3为酒店。[{"businessId":11925099374,"businessType":3,"displayItemInNonWorkingTime":1,"mondayConfirmLocalTime":"09:00-18:00","operator":"操作人","orderConfirmType":1,"saturdayConfirmLocalTime":"09:00-18:00","sellerId":2054718374,"sellerNick":"sandbox_b_27","sundayConfirmLocalTime":"09:00-18:00","supplier":"","thursdayConfirmLocalTime":"09:00-18:00","timeZoneName":"Asia/Shanghai","tuesdayConfirmLocalTime":"09:00-18:00","wednesdayConfirmLocalTime":"09:00-18:00","fridayConfirmLocalTime":"09:00-18:00"},{"businessId":11925099374,"businessType":3,"displayItemInNonWorkingTime":1,"mondayConfirmLocalTime":"09:00-18:00","operator":"操作人","orderConfirmType":2,"saturdayConfirmLocalTime":"09:00-18:00","sellerId":2054718374,"sellerNick":"sandbox_b_27","sundayConfirmLocalTime":"09:00-18:00","supplier":"","thursdayConfirmLocalTime":"09:00-18:00","timeZoneName":"Asia/Shanghai","tuesdayConfirmLocalTime":"09:00-18:00","wednesdayConfirmLocalTime":"09:00-18:00","fridayConfirmLocalTime":"09:00-18:00"}]
func (r *TaobaoXhotelServicetimeUpdateAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetParam() string {
	return r._param
}
