package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableAddAPIResponse 智能应用服务登记工作表数据新增 API返回值
// taobao.smartapp.table.add
//
// 智能应用服务登记工作表数据新增
type TaobaoSmartappTableAddAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappTableAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartappTableAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartappTableAddAPIResponseModel).Reset()
}

// TaobaoSmartappTableAddAPIResponseModel is 智能应用服务登记工作表数据新增 成功返回结果
type TaobaoSmartappTableAddAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单条新增记录ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 错误提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSmartappTableAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Id = ""
	m.ErrorMsg = ""
	m.BusinessSuccess = false
}

var poolTaobaoSmartappTableAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartappTableAddAPIResponse)
	},
}

// GetTaobaoSmartappTableAddAPIResponse 从 sync.Pool 获取 TaobaoSmartappTableAddAPIResponse
func GetTaobaoSmartappTableAddAPIResponse() *TaobaoSmartappTableAddAPIResponse {
	return poolTaobaoSmartappTableAddAPIResponse.Get().(*TaobaoSmartappTableAddAPIResponse)
}

// ReleaseTaobaoSmartappTableAddAPIResponse 将 TaobaoSmartappTableAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartappTableAddAPIResponse(v *TaobaoSmartappTableAddAPIResponse) {
	v.Reset()
	poolTaobaoSmartappTableAddAPIResponse.Put(v)
}
