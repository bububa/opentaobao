package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrInventoryInitialAPIResponse 门店库存初始化前后端商品绑定 API返回值
// tmall.nr.inventory.initial
//
// 用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
type TmallNrInventoryInitialAPIResponse struct {
	model.CommonResponse
	TmallNrInventoryInitialAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrInventoryInitialAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrInventoryInitialAPIResponseModel).Reset()
}

// TmallNrInventoryInitialAPIResponseModel is 门店库存初始化前后端商品绑定 成功返回结果
type TmallNrInventoryInitialAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_inventory_initial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息提示
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误编码
	ErrorCode2 string `json:"error_code2,omitempty" xml:"error_code2,omitempty"`
	// 初始化成功
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrInventoryInitialAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessage = ""
	m.ErrorCode2 = ""
	m.ResultData = false
	m.IsSuccess = false
}

var poolTmallNrInventoryInitialAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrInventoryInitialAPIResponse)
	},
}

// GetTmallNrInventoryInitialAPIResponse 从 sync.Pool 获取 TmallNrInventoryInitialAPIResponse
func GetTmallNrInventoryInitialAPIResponse() *TmallNrInventoryInitialAPIResponse {
	return poolTmallNrInventoryInitialAPIResponse.Get().(*TmallNrInventoryInitialAPIResponse)
}

// ReleaseTmallNrInventoryInitialAPIResponse 将 TmallNrInventoryInitialAPIResponse 保存到 sync.Pool
func ReleaseTmallNrInventoryInitialAPIResponse(v *TmallNrInventoryInitialAPIResponse) {
	v.Reset()
	poolTmallNrInventoryInitialAPIResponse.Put(v)
}
