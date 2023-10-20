package usergrowth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthAdMaterialUpdateAPIResponse 素材更新 API返回值
// taobao.usergrowth.ad.material.update
//
// 素材更新
type TaobaoUsergrowthAdMaterialUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthAdMaterialUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsergrowthAdMaterialUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsergrowthAdMaterialUpdateAPIResponseModel).Reset()
}

// TaobaoUsergrowthAdMaterialUpdateAPIResponseModel is 素材更新 成功返回结果
type TaobaoUsergrowthAdMaterialUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_ad_material_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 更新创意的id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 参数错误
	ResponseCode int64 `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 请求结果
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsergrowthAdMaterialUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.Data = ""
	m.Message = ""
	m.ResponseCode = 0
	m.Successful = false
}

var poolTaobaoUsergrowthAdMaterialUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsergrowthAdMaterialUpdateAPIResponse)
	},
}

// GetTaobaoUsergrowthAdMaterialUpdateAPIResponse 从 sync.Pool 获取 TaobaoUsergrowthAdMaterialUpdateAPIResponse
func GetTaobaoUsergrowthAdMaterialUpdateAPIResponse() *TaobaoUsergrowthAdMaterialUpdateAPIResponse {
	return poolTaobaoUsergrowthAdMaterialUpdateAPIResponse.Get().(*TaobaoUsergrowthAdMaterialUpdateAPIResponse)
}

// ReleaseTaobaoUsergrowthAdMaterialUpdateAPIResponse 将 TaobaoUsergrowthAdMaterialUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsergrowthAdMaterialUpdateAPIResponse(v *TaobaoUsergrowthAdMaterialUpdateAPIResponse) {
	v.Reset()
	poolTaobaoUsergrowthAdMaterialUpdateAPIResponse.Put(v)
}
