package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrSellerStorerangeReadAPIResponse 门店服务范围读取 API返回值
// tmall.nr.seller.storerange.read
//
// 读取卖家所属门店的服务范围
type TmallNrSellerStorerangeReadAPIResponse struct {
	model.CommonResponse
	TmallNrSellerStorerangeReadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrSellerStorerangeReadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrSellerStorerangeReadAPIResponseModel).Reset()
}

// TmallNrSellerStorerangeReadAPIResponseModel is 门店服务范围读取 成功返回结果
type TmallNrSellerStorerangeReadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_seller_storerange_read_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	ResultList *NrResult `json:"result_list,omitempty" xml:"result_list,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrSellerStorerangeReadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = nil
}

var poolTmallNrSellerStorerangeReadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrSellerStorerangeReadAPIResponse)
	},
}

// GetTmallNrSellerStorerangeReadAPIResponse 从 sync.Pool 获取 TmallNrSellerStorerangeReadAPIResponse
func GetTmallNrSellerStorerangeReadAPIResponse() *TmallNrSellerStorerangeReadAPIResponse {
	return poolTmallNrSellerStorerangeReadAPIResponse.Get().(*TmallNrSellerStorerangeReadAPIResponse)
}

// ReleaseTmallNrSellerStorerangeReadAPIResponse 将 TmallNrSellerStorerangeReadAPIResponse 保存到 sync.Pool
func ReleaseTmallNrSellerStorerangeReadAPIResponse(v *TmallNrSellerStorerangeReadAPIResponse) {
	v.Reset()
	poolTmallNrSellerStorerangeReadAPIResponse.Put(v)
}
