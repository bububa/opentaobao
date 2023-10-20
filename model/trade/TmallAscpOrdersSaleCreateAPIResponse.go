package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAscpOrdersSaleCreateAPIResponse ASCP渠道中心销售单创建接口 API返回值
// tmall.ascp.orders.sale.create
//
// ASCP渠道中心销售单创建接口
type TmallAscpOrdersSaleCreateAPIResponse struct {
	model.CommonResponse
	TmallAscpOrdersSaleCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAscpOrdersSaleCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAscpOrdersSaleCreateAPIResponseModel).Reset()
}

// TmallAscpOrdersSaleCreateAPIResponseModel is ASCP渠道中心销售单创建接口 成功返回结果
type TmallAscpOrdersSaleCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ascp_orders_sale_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallAscpOrdersSaleCreateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAscpOrdersSaleCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAscpOrdersSaleCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAscpOrdersSaleCreateAPIResponse)
	},
}

// GetTmallAscpOrdersSaleCreateAPIResponse 从 sync.Pool 获取 TmallAscpOrdersSaleCreateAPIResponse
func GetTmallAscpOrdersSaleCreateAPIResponse() *TmallAscpOrdersSaleCreateAPIResponse {
	return poolTmallAscpOrdersSaleCreateAPIResponse.Get().(*TmallAscpOrdersSaleCreateAPIResponse)
}

// ReleaseTmallAscpOrdersSaleCreateAPIResponse 将 TmallAscpOrdersSaleCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallAscpOrdersSaleCreateAPIResponse(v *TmallAscpOrdersSaleCreateAPIResponse) {
	v.Reset()
	poolTmallAscpOrdersSaleCreateAPIResponse.Put(v)
}
