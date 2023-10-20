package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrOrderLogisInfoAPIResponse 区域零售订单获取取件码 API返回值
// tmall.nr.order.logis.info
//
// 区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
type TmallNrOrderLogisInfoAPIResponse struct {
	model.CommonResponse
	TmallNrOrderLogisInfoAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrOrderLogisInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrOrderLogisInfoAPIResponseModel).Reset()
}

// TmallNrOrderLogisInfoAPIResponseModel is 区域零售订单获取取件码 成功返回结果
type TmallNrOrderLogisInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_order_logis_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果实体
	Result *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrOrderLogisInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrOrderLogisInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrOrderLogisInfoAPIResponse)
	},
}

// GetTmallNrOrderLogisInfoAPIResponse 从 sync.Pool 获取 TmallNrOrderLogisInfoAPIResponse
func GetTmallNrOrderLogisInfoAPIResponse() *TmallNrOrderLogisInfoAPIResponse {
	return poolTmallNrOrderLogisInfoAPIResponse.Get().(*TmallNrOrderLogisInfoAPIResponse)
}

// ReleaseTmallNrOrderLogisInfoAPIResponse 将 TmallNrOrderLogisInfoAPIResponse 保存到 sync.Pool
func ReleaseTmallNrOrderLogisInfoAPIResponse(v *TmallNrOrderLogisInfoAPIResponse) {
	v.Reset()
	poolTmallNrOrderLogisInfoAPIResponse.Put(v)
}
