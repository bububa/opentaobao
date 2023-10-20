package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasServiceValidateAPIResponse 增值服务订购服务验证 API返回值
// taobao.vas.service.validate
//
// 增值服务订购服务验证
type TaobaoVasServiceValidateAPIResponse struct {
	model.CommonResponse
	TaobaoVasServiceValidateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVasServiceValidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVasServiceValidateAPIResponseModel).Reset()
}

// TaobaoVasServiceValidateAPIResponseModel is 增值服务订购服务验证 成功返回结果
type TaobaoVasServiceValidateAPIResponseModel struct {
	XMLName xml.Name `xml:"vas_service_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// none 没有订购,open 已经开通服务,freeze 服务已经冻结,close 服务已经关闭,error 系统错误
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVasServiceValidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Status = ""
}

var poolTaobaoVasServiceValidateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVasServiceValidateAPIResponse)
	},
}

// GetTaobaoVasServiceValidateAPIResponse 从 sync.Pool 获取 TaobaoVasServiceValidateAPIResponse
func GetTaobaoVasServiceValidateAPIResponse() *TaobaoVasServiceValidateAPIResponse {
	return poolTaobaoVasServiceValidateAPIResponse.Get().(*TaobaoVasServiceValidateAPIResponse)
}

// ReleaseTaobaoVasServiceValidateAPIResponse 将 TaobaoVasServiceValidateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVasServiceValidateAPIResponse(v *TaobaoVasServiceValidateAPIResponse) {
	v.Reset()
	poolTaobaoVasServiceValidateAPIResponse.Put(v)
}
