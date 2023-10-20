package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenCombineitemSynchronizeAPIResponse 组合商品接口 API返回值
// taobao.qimen.combineitem.synchronize
//
// ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoQimenCombineitemSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenCombineitemSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenCombineitemSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenCombineitemSynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenCombineitemSynchronizeAPIResponseModel is 组合商品接口 成功返回结果
type TaobaoQimenCombineitemSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_combineitem_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenCombineitemSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenCombineitemSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenCombineitemSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemSynchronizeAPIResponse)
	},
}

// GetTaobaoQimenCombineitemSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenCombineitemSynchronizeAPIResponse
func GetTaobaoQimenCombineitemSynchronizeAPIResponse() *TaobaoQimenCombineitemSynchronizeAPIResponse {
	return poolTaobaoQimenCombineitemSynchronizeAPIResponse.Get().(*TaobaoQimenCombineitemSynchronizeAPIResponse)
}

// ReleaseTaobaoQimenCombineitemSynchronizeAPIResponse 将 TaobaoQimenCombineitemSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenCombineitemSynchronizeAPIResponse(v *TaobaoQimenCombineitemSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenCombineitemSynchronizeAPIResponse.Put(v)
}
