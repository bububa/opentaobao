package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSingleitemSynchronizeAPIResponse 商品同步接口 API返回值
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
type TaobaoQimenSingleitemSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenSingleitemSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenSingleitemSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenSingleitemSynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenSingleitemSynchronizeAPIResponseModel is 商品同步接口 成功返回结果
type TaobaoQimenSingleitemSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_singleitem_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenSingleitemSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenSingleitemSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenSingleitemSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSingleitemSynchronizeAPIResponse)
	},
}

// GetTaobaoQimenSingleitemSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenSingleitemSynchronizeAPIResponse
func GetTaobaoQimenSingleitemSynchronizeAPIResponse() *TaobaoQimenSingleitemSynchronizeAPIResponse {
	return poolTaobaoQimenSingleitemSynchronizeAPIResponse.Get().(*TaobaoQimenSingleitemSynchronizeAPIResponse)
}

// ReleaseTaobaoQimenSingleitemSynchronizeAPIResponse 将 TaobaoQimenSingleitemSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenSingleitemSynchronizeAPIResponse(v *TaobaoQimenSingleitemSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenSingleitemSynchronizeAPIResponse.Put(v)
}
