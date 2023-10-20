package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMemberpriceUpdateAPIResponse 商品售价会员价修改 API返回值
// alibaba.wdk.item.memberprice.update
//
// 商品售价会员价修改
type AlibabaWdkItemMemberpriceUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMemberpriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemMemberpriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemMemberpriceUpdateAPIResponseModel).Reset()
}

// AlibabaWdkItemMemberpriceUpdateAPIResponseModel is 商品售价会员价修改 成功返回结果
type AlibabaWdkItemMemberpriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_memberprice_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMemberpriceUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemMemberpriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemMemberpriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMemberpriceUpdateAPIResponse)
	},
}

// GetAlibabaWdkItemMemberpriceUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemMemberpriceUpdateAPIResponse
func GetAlibabaWdkItemMemberpriceUpdateAPIResponse() *AlibabaWdkItemMemberpriceUpdateAPIResponse {
	return poolAlibabaWdkItemMemberpriceUpdateAPIResponse.Get().(*AlibabaWdkItemMemberpriceUpdateAPIResponse)
}

// ReleaseAlibabaWdkItemMemberpriceUpdateAPIResponse 将 AlibabaWdkItemMemberpriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemMemberpriceUpdateAPIResponse(v *AlibabaWdkItemMemberpriceUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemMemberpriceUpdateAPIResponse.Put(v)
}
