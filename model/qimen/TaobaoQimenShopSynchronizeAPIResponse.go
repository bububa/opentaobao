package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenShopSynchronizeAPIResponse 店铺同步接口 API返回值
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
type TaobaoQimenShopSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenShopSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenShopSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenShopSynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenShopSynchronizeAPIResponseModel is 店铺同步接口 成功返回结果
type TaobaoQimenShopSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_shop_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Response
	Response *TaobaoQimenShopSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenShopSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenShopSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenShopSynchronizeAPIResponse)
	},
}

// GetTaobaoQimenShopSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenShopSynchronizeAPIResponse
func GetTaobaoQimenShopSynchronizeAPIResponse() *TaobaoQimenShopSynchronizeAPIResponse {
	return poolTaobaoQimenShopSynchronizeAPIResponse.Get().(*TaobaoQimenShopSynchronizeAPIResponse)
}

// ReleaseTaobaoQimenShopSynchronizeAPIResponse 将 TaobaoQimenShopSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenShopSynchronizeAPIResponse(v *TaobaoQimenShopSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenShopSynchronizeAPIResponse.Put(v)
}
