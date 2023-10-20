package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkMealPosGetfetchmealcodeAPIResponse 五道口餐饮取餐号获取接口 API返回值
// wdk.meal.pos.getfetchmealcode
//
// pos机创建订单前获取餐饮取餐号
type WdkMealPosGetfetchmealcodeAPIResponse struct {
	model.CommonResponse
	WdkMealPosGetfetchmealcodeAPIResponseModel
}

// Reset 清空结构体
func (m *WdkMealPosGetfetchmealcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkMealPosGetfetchmealcodeAPIResponseModel).Reset()
}

// WdkMealPosGetfetchmealcodeAPIResponseModel is 五道口餐饮取餐号获取接口 成功返回结果
type WdkMealPosGetfetchmealcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_meal_pos_getfetchmealcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取餐号
	FetchMealCode string `json:"fetch_meal_code,omitempty" xml:"fetch_meal_code,omitempty"`
}

// Reset 清空结构体
func (m *WdkMealPosGetfetchmealcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FetchMealCode = ""
}

var poolWdkMealPosGetfetchmealcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkMealPosGetfetchmealcodeAPIResponse)
	},
}

// GetWdkMealPosGetfetchmealcodeAPIResponse 从 sync.Pool 获取 WdkMealPosGetfetchmealcodeAPIResponse
func GetWdkMealPosGetfetchmealcodeAPIResponse() *WdkMealPosGetfetchmealcodeAPIResponse {
	return poolWdkMealPosGetfetchmealcodeAPIResponse.Get().(*WdkMealPosGetfetchmealcodeAPIResponse)
}

// ReleaseWdkMealPosGetfetchmealcodeAPIResponse 将 WdkMealPosGetfetchmealcodeAPIResponse 保存到 sync.Pool
func ReleaseWdkMealPosGetfetchmealcodeAPIResponse(v *WdkMealPosGetfetchmealcodeAPIResponse) {
	v.Reset()
	poolWdkMealPosGetfetchmealcodeAPIResponse.Put(v)
}
