package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse 库存初始化接口 API返回值
// taobao.jst.astrolabe.storeinventory.iteminitial
//
// ERP调用奇门的接口，对门店的库存进行初始化
type TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel is 库存初始化接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_iteminitial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息列表
	ErrorDescriptions []TaobaoJstAstrolabeStoreinventoryIteminitialError `json:"error_descriptions,omitempty" xml:"error_descriptions>taobao_jst_astrolabe_storeinventory_iteminitial_error,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应标签
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorDescriptions = m.ErrorDescriptions[:0]
	m.Flag = ""
	m.QimenCode = ""
	m.Message = ""
}

var poolTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse
func GetTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse() *TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse {
	return poolTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse.Get().(*TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse 将 TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse(v *TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse.Put(v)
}
