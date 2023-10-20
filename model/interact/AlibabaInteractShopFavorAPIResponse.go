package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractShopFavorAPIResponse 店铺收藏 API返回值
// alibaba.interact.shop.favor
//
// 店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
type AlibabaInteractShopFavorAPIResponse struct {
	model.CommonResponse
	AlibabaInteractShopFavorAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractShopFavorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractShopFavorAPIResponseModel).Reset()
}

// AlibabaInteractShopFavorAPIResponseModel is 店铺收藏 成功返回结果
type AlibabaInteractShopFavorAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_shop_favor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractShopFavorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractShopFavorAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractShopFavorAPIResponse)
	},
}

// GetAlibabaInteractShopFavorAPIResponse 从 sync.Pool 获取 AlibabaInteractShopFavorAPIResponse
func GetAlibabaInteractShopFavorAPIResponse() *AlibabaInteractShopFavorAPIResponse {
	return poolAlibabaInteractShopFavorAPIResponse.Get().(*AlibabaInteractShopFavorAPIResponse)
}

// ReleaseAlibabaInteractShopFavorAPIResponse 将 AlibabaInteractShopFavorAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractShopFavorAPIResponse(v *AlibabaInteractShopFavorAPIResponse) {
	v.Reset()
	poolAlibabaInteractShopFavorAPIResponse.Put(v)
}
