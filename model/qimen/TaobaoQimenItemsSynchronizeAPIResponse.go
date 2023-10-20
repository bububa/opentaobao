package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemsSynchronizeAPIResponse 商品同步接口 (批量) API返回值
// taobao.qimen.items.synchronize
//
// ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoQimenItemsSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemsSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenItemsSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenItemsSynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenItemsSynchronizeAPIResponseModel is 商品同步接口 (批量) 成功返回结果
type TaobaoQimenItemsSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_items_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ItemsSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenItemsSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenItemsSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemsSynchronizeAPIResponse)
	},
}

// GetTaobaoQimenItemsSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenItemsSynchronizeAPIResponse
func GetTaobaoQimenItemsSynchronizeAPIResponse() *TaobaoQimenItemsSynchronizeAPIResponse {
	return poolTaobaoQimenItemsSynchronizeAPIResponse.Get().(*TaobaoQimenItemsSynchronizeAPIResponse)
}

// ReleaseTaobaoQimenItemsSynchronizeAPIResponse 将 TaobaoQimenItemsSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenItemsSynchronizeAPIResponse(v *TaobaoQimenItemsSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenItemsSynchronizeAPIResponse.Put(v)
}
