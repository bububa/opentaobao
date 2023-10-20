package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableListGetAPIResponse 智能应用服务登记工作表列表查询 API返回值
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
type TaobaoSmartappTableListGetAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappTableListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartappTableListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartappTableListGetAPIResponseModel).Reset()
}

// TaobaoSmartappTableListGetAPIResponseModel is 智能应用服务登记工作表列表查询 成功返回结果
type TaobaoSmartappTableListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用结果
	Result *AfterSaleGetWorkTableListResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSmartappTableListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = nil
	m.BusinessSuccess = false
}

var poolTaobaoSmartappTableListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartappTableListGetAPIResponse)
	},
}

// GetTaobaoSmartappTableListGetAPIResponse 从 sync.Pool 获取 TaobaoSmartappTableListGetAPIResponse
func GetTaobaoSmartappTableListGetAPIResponse() *TaobaoSmartappTableListGetAPIResponse {
	return poolTaobaoSmartappTableListGetAPIResponse.Get().(*TaobaoSmartappTableListGetAPIResponse)
}

// ReleaseTaobaoSmartappTableListGetAPIResponse 将 TaobaoSmartappTableListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartappTableListGetAPIResponse(v *TaobaoSmartappTableListGetAPIResponse) {
	v.Reset()
	poolTaobaoSmartappTableListGetAPIResponse.Put(v)
}
