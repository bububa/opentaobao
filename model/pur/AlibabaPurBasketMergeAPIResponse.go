package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurBasketMergeAPIResponse 合并购物车 API返回值
// alibaba.pur.basket.merge
//
// 采购商城接入第三方商家合并购物车接口服务
type AlibabaPurBasketMergeAPIResponse struct {
	model.CommonResponse
	AlibabaPurBasketMergeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurBasketMergeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurBasketMergeAPIResponseModel).Reset()
}

// AlibabaPurBasketMergeAPIResponseModel is 合并购物车 成功返回结果
type AlibabaPurBasketMergeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_basket_merge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurBasketMergeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurBasketMergeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurBasketMergeAPIResponse)
	},
}

// GetAlibabaPurBasketMergeAPIResponse 从 sync.Pool 获取 AlibabaPurBasketMergeAPIResponse
func GetAlibabaPurBasketMergeAPIResponse() *AlibabaPurBasketMergeAPIResponse {
	return poolAlibabaPurBasketMergeAPIResponse.Get().(*AlibabaPurBasketMergeAPIResponse)
}

// ReleaseAlibabaPurBasketMergeAPIResponse 将 AlibabaPurBasketMergeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurBasketMergeAPIResponse(v *AlibabaPurBasketMergeAPIResponse) {
	v.Reset()
	poolAlibabaPurBasketMergeAPIResponse.Put(v)
}
