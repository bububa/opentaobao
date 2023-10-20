package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductReleaseAPIResponse 供应商铺货 API返回值
// tmall.supplychain.channel.product.release
//
// 供应商渠道铺货接口
type TmallSupplychainChannelProductReleaseAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductReleaseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductReleaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductReleaseAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductReleaseAPIResponseModel is 供应商铺货 成功返回结果
type TmallSupplychainChannelProductReleaseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_release_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallSupplychainChannelProductReleaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductReleaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductReleaseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductReleaseAPIResponse)
	},
}

// GetTmallSupplychainChannelProductReleaseAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductReleaseAPIResponse
func GetTmallSupplychainChannelProductReleaseAPIResponse() *TmallSupplychainChannelProductReleaseAPIResponse {
	return poolTmallSupplychainChannelProductReleaseAPIResponse.Get().(*TmallSupplychainChannelProductReleaseAPIResponse)
}

// ReleaseTmallSupplychainChannelProductReleaseAPIResponse 将 TmallSupplychainChannelProductReleaseAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductReleaseAPIResponse(v *TmallSupplychainChannelProductReleaseAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductReleaseAPIResponse.Put(v)
}
