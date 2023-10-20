package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse 库存增量更新接口 API返回值
// taobao.jst.astrolabe.storeinventory.itemupdate
//
// ERP调用该接口，增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存。
type TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponseModel).Reset()
}

// TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponseModel is 库存增量更新接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_itemupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息列表
	ErrorDescriptions []TaobaoJstAstrolabeStoreinventoryItemupdateError `json:"error_descriptions,omitempty" xml:"error_descriptions>taobao_jst_astrolabe_storeinventory_itemupdate_error,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应标签
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorDescriptions = m.ErrorDescriptions[:0]
	m.Flag = ""
	m.QimenCode = ""
	m.Message = ""
}

var poolTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse 从 sync.Pool 获取 TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse
func GetTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse() *TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse {
	return poolTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse.Get().(*TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse 将 TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse(v *TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse) {
	v.Reset()
	poolTaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse.Put(v)
}
