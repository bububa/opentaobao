package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryruleCreateAPIResponse 渠道间库存规则设置接口 API返回值
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
type TaobaoQimenInventoryruleCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventoryruleCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryruleCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventoryruleCreateAPIResponseModel).Reset()
}

// TaobaoQimenInventoryruleCreateAPIResponseModel is 渠道间库存规则设置接口 成功返回结果
type TaobaoQimenInventoryruleCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventoryrule_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryruleCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventoryruleCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryruleCreateAPIResponse)
	},
}

// GetTaobaoQimenInventoryruleCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenInventoryruleCreateAPIResponse
func GetTaobaoQimenInventoryruleCreateAPIResponse() *TaobaoQimenInventoryruleCreateAPIResponse {
	return poolTaobaoQimenInventoryruleCreateAPIResponse.Get().(*TaobaoQimenInventoryruleCreateAPIResponse)
}

// ReleaseTaobaoQimenInventoryruleCreateAPIResponse 将 TaobaoQimenInventoryruleCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventoryruleCreateAPIResponse(v *TaobaoQimenInventoryruleCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventoryruleCreateAPIResponse.Put(v)
}
