package xhotelitem

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelServicetimeUpdateAPIRequest) Reset() {
	r._param = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.servicetime.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请按照示例值的格式来填写，涉及到是否当日订单，是否展示，周一到周日的服务时间，业务id,业务类型1为卖家，2为supplier ，3为酒店。[{&#34;businessId&#34;:11925099374,&#34;businessType&#34;:3,&#34;displayItemInNonWorkingTime&#34;:1,&#34;mondayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;operator&#34;:&#34;操作人&#34;,&#34;orderConfirmType&#34;:1,&#34;saturdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;sellerId&#34;:2054718374,&#34;sellerNick&#34;:&#34;sandbox_b_27&#34;,&#34;sundayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;supplier&#34;:&#34;&#34;,&#34;thursdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;timeZoneName&#34;:&#34;Asia/Shanghai&#34;,&#34;tuesdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;wednesdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;fridayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;},{&#34;businessId&#34;:11925099374,&#34;businessType&#34;:3,&#34;displayItemInNonWorkingTime&#34;:1,&#34;mondayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;operator&#34;:&#34;操作人&#34;,&#34;orderConfirmType&#34;:2,&#34;saturdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;sellerId&#34;:2054718374,&#34;sellerNick&#34;:&#34;sandbox_b_27&#34;,&#34;sundayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;supplier&#34;:&#34;&#34;,&#34;thursdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;timeZoneName&#34;:&#34;Asia/Shanghai&#34;,&#34;tuesdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;wednesdayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;,&#34;fridayConfirmLocalTime&#34;:&#34;09:00-18:00&#34;}]
func (r *TaobaoXhotelServicetimeUpdateAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoXhotelServicetimeUpdateAPIRequest) GetParam() string {
	return r._param
}

var poolTaobaoXhotelServicetimeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelServicetimeUpdateRequest()
	},
}

// GetTaobaoXhotelServicetimeUpdateRequest 从 sync.Pool 获取 TaobaoXhotelServicetimeUpdateAPIRequest
func GetTaobaoXhotelServicetimeUpdateAPIRequest() *TaobaoXhotelServicetimeUpdateAPIRequest {
	return poolTaobaoXhotelServicetimeUpdateAPIRequest.Get().(*TaobaoXhotelServicetimeUpdateAPIRequest)
}

// ReleaseTaobaoXhotelServicetimeUpdateAPIRequest 将 TaobaoXhotelServicetimeUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelServicetimeUpdateAPIRequest(v *TaobaoXhotelServicetimeUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelServicetimeUpdateAPIRequest.Put(v)
}
