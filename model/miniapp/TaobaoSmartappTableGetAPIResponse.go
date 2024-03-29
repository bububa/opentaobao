package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableGetAPIResponse 智能应用服务登记工作表数据查询 API返回值
// taobao.smartapp.table.get
//
// 智能应用服务登记工作表数据查询
type TaobaoSmartappTableGetAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappTableGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartappTableGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartappTableGetAPIResponseModel).Reset()
}

// TaobaoSmartappTableGetAPIResponseModel is 智能应用服务登记工作表数据查询 成功返回结果
type TaobaoSmartappTableGetAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 查询结果
	Result *AfterSaleTableSelectResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSmartappTableGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = nil
	m.BusinessSuccess = false
}

var poolTaobaoSmartappTableGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartappTableGetAPIResponse)
	},
}

// GetTaobaoSmartappTableGetAPIResponse 从 sync.Pool 获取 TaobaoSmartappTableGetAPIResponse
func GetTaobaoSmartappTableGetAPIResponse() *TaobaoSmartappTableGetAPIResponse {
	return poolTaobaoSmartappTableGetAPIResponse.Get().(*TaobaoSmartappTableGetAPIResponse)
}

// ReleaseTaobaoSmartappTableGetAPIResponse 将 TaobaoSmartappTableGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartappTableGetAPIResponse(v *TaobaoSmartappTableGetAPIResponse) {
	v.Reset()
	poolTaobaoSmartappTableGetAPIResponse.Put(v)
}
