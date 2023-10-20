package dengta

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaOrderEffectImportAPIResponse 天下秀订单数据导入 API返回值
// alibaba.pictures.dengta.order.effect.import
//
// 提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
type AlibabaPicturesDengtaOrderEffectImportAPIResponse struct {
	model.CommonResponse
	AlibabaPicturesDengtaOrderEffectImportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaOrderEffectImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPicturesDengtaOrderEffectImportAPIResponseModel).Reset()
}

// AlibabaPicturesDengtaOrderEffectImportAPIResponseModel is 天下秀订单数据导入 成功返回结果
type AlibabaPicturesDengtaOrderEffectImportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pictures_dengta_order_effect_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPicturesDengtaOrderEffectImportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPicturesDengtaOrderEffectImportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPicturesDengtaOrderEffectImportAPIResponse)
	},
}

// GetAlibabaPicturesDengtaOrderEffectImportAPIResponse 从 sync.Pool 获取 AlibabaPicturesDengtaOrderEffectImportAPIResponse
func GetAlibabaPicturesDengtaOrderEffectImportAPIResponse() *AlibabaPicturesDengtaOrderEffectImportAPIResponse {
	return poolAlibabaPicturesDengtaOrderEffectImportAPIResponse.Get().(*AlibabaPicturesDengtaOrderEffectImportAPIResponse)
}

// ReleaseAlibabaPicturesDengtaOrderEffectImportAPIResponse 将 AlibabaPicturesDengtaOrderEffectImportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPicturesDengtaOrderEffectImportAPIResponse(v *AlibabaPicturesDengtaOrderEffectImportAPIResponse) {
	v.Reset()
	poolAlibabaPicturesDengtaOrderEffectImportAPIResponse.Put(v)
}
