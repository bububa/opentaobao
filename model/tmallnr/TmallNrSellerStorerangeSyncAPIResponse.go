package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrSellerStorerangeSyncAPIResponse 同步商户中心服务范围 API返回值
// tmall.nr.seller.storerange.sync
//
// 同步商户中心服务范围
type TmallNrSellerStorerangeSyncAPIResponse struct {
	model.CommonResponse
	TmallNrSellerStorerangeSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrSellerStorerangeSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrSellerStorerangeSyncAPIResponseModel).Reset()
}

// TmallNrSellerStorerangeSyncAPIResponseModel is 同步商户中心服务范围 成功返回结果
type TmallNrSellerStorerangeSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_seller_storerange_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 请求是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrSellerStorerangeSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultData = false
	m.SuccessFlag = false
}

var poolTmallNrSellerStorerangeSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrSellerStorerangeSyncAPIResponse)
	},
}

// GetTmallNrSellerStorerangeSyncAPIResponse 从 sync.Pool 获取 TmallNrSellerStorerangeSyncAPIResponse
func GetTmallNrSellerStorerangeSyncAPIResponse() *TmallNrSellerStorerangeSyncAPIResponse {
	return poolTmallNrSellerStorerangeSyncAPIResponse.Get().(*TmallNrSellerStorerangeSyncAPIResponse)
}

// ReleaseTmallNrSellerStorerangeSyncAPIResponse 将 TmallNrSellerStorerangeSyncAPIResponse 保存到 sync.Pool
func ReleaseTmallNrSellerStorerangeSyncAPIResponse(v *TmallNrSellerStorerangeSyncAPIResponse) {
	v.Reset()
	poolTmallNrSellerStorerangeSyncAPIResponse.Put(v)
}
