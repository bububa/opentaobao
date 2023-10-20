package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductReleaseStatusGetAPIResponse 产品铺货状态查询 API返回值
// tmall.supplychain.channel.product.release.status.get
//
// 巴拿马战役渠道产品状态查询
type TmallSupplychainChannelProductReleaseStatusGetAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductReleaseStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductReleaseStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductReleaseStatusGetAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductReleaseStatusGetAPIResponseModel is 产品铺货状态查询 成功返回结果
type TmallSupplychainChannelProductReleaseStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_release_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallSupplychainChannelProductReleaseStatusGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductReleaseStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductReleaseStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductReleaseStatusGetAPIResponse)
	},
}

// GetTmallSupplychainChannelProductReleaseStatusGetAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductReleaseStatusGetAPIResponse
func GetTmallSupplychainChannelProductReleaseStatusGetAPIResponse() *TmallSupplychainChannelProductReleaseStatusGetAPIResponse {
	return poolTmallSupplychainChannelProductReleaseStatusGetAPIResponse.Get().(*TmallSupplychainChannelProductReleaseStatusGetAPIResponse)
}

// ReleaseTmallSupplychainChannelProductReleaseStatusGetAPIResponse 将 TmallSupplychainChannelProductReleaseStatusGetAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductReleaseStatusGetAPIResponse(v *TmallSupplychainChannelProductReleaseStatusGetAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductReleaseStatusGetAPIResponse.Put(v)
}
