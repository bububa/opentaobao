package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryInitialAPIResponse 后端商品库存初始化 API返回值
// taobao.jst.astrolabe.storeinventory.initial
//
// 初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。
type TaobaoJstAstrolabeStoreinventoryInitialAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryInitialAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryInitialAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeStoreinventoryInitialAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeStoreinventoryInitialAPIResponseModel is 后端商品库存初始化 成功返回结果
type TaobaoJstAstrolabeStoreinventoryInitialAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_initial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息列表
	ErrorDescriptions []TaobaoJstAstrolabeStoreinventoryInitialError `json:"error_descriptions,omitempty" xml:"error_descriptions>taobao_jst_astrolabe_storeinventory_initial_error,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应标签
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryInitialAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorDescriptions = m.ErrorDescriptions[:0]
	m.Message = ""
	m.ResultCode = ""
	m.Flag = ""
}

var poolTaobaoJstAstrolabeStoreinventoryInitialAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryInitialAPIResponse)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryInitialAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryInitialAPIResponse
func GetTaobaoJstAstrolabeStoreinventoryInitialAPIResponse() *TaobaoJstAstrolabeStoreinventoryInitialAPIResponse {
	return poolTaobaoJstAstrolabeStoreinventoryInitialAPIResponse.Get().(*TaobaoJstAstrolabeStoreinventoryInitialAPIResponse)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryInitialAPIResponse 将 TaobaoJstAstrolabeStoreinventoryInitialAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryInitialAPIResponse(v *TaobaoJstAstrolabeStoreinventoryInitialAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryInitialAPIResponse.Put(v)
}
