package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStorecategoryGetAPIResponse 门店类目获取接口 API返回值
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
type TaobaoQimenStorecategoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStorecategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStorecategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStorecategoryGetAPIResponseModel).Reset()
}

// TaobaoQimenStorecategoryGetAPIResponseModel is 门店类目获取接口 成功返回结果
type TaobaoQimenStorecategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_storecategory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应code
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 类目json字符串
	StoreCategory string `json:"store_category,omitempty" xml:"store_category,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStorecategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Flag = ""
	m.QimenCode = ""
	m.StoreCategory = ""
}

var poolTaobaoQimenStorecategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStorecategoryGetAPIResponse)
	},
}

// GetTaobaoQimenStorecategoryGetAPIResponse 从 sync.Pool 获取 TaobaoQimenStorecategoryGetAPIResponse
func GetTaobaoQimenStorecategoryGetAPIResponse() *TaobaoQimenStorecategoryGetAPIResponse {
	return poolTaobaoQimenStorecategoryGetAPIResponse.Get().(*TaobaoQimenStorecategoryGetAPIResponse)
}

// ReleaseTaobaoQimenStorecategoryGetAPIResponse 将 TaobaoQimenStorecategoryGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStorecategoryGetAPIResponse(v *TaobaoQimenStorecategoryGetAPIResponse) {
	v.Reset()
	poolTaobaoQimenStorecategoryGetAPIResponse.Put(v)
}
