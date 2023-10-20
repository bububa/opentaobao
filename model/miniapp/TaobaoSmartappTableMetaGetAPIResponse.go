package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableMetaGetAPIResponse 智能应用服务登记工作表元数据查询 API返回值
// taobao.smartapp.table.meta.get
//
// 智能应用服务登记工作表元数据查询
type TaobaoSmartappTableMetaGetAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappTableMetaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartappTableMetaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartappTableMetaGetAPIResponseModel).Reset()
}

// TaobaoSmartappTableMetaGetAPIResponseModel is 智能应用服务登记工作表元数据查询 成功返回结果
type TaobaoSmartappTableMetaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_meta_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Result *AfterSaleFieldMetaResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	BusinessSuccess bool `json:"business_success,omitempty" xml:"business_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSmartappTableMetaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = nil
	m.BusinessSuccess = false
}

var poolTaobaoSmartappTableMetaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartappTableMetaGetAPIResponse)
	},
}

// GetTaobaoSmartappTableMetaGetAPIResponse 从 sync.Pool 获取 TaobaoSmartappTableMetaGetAPIResponse
func GetTaobaoSmartappTableMetaGetAPIResponse() *TaobaoSmartappTableMetaGetAPIResponse {
	return poolTaobaoSmartappTableMetaGetAPIResponse.Get().(*TaobaoSmartappTableMetaGetAPIResponse)
}

// ReleaseTaobaoSmartappTableMetaGetAPIResponse 将 TaobaoSmartappTableMetaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartappTableMetaGetAPIResponse(v *TaobaoSmartappTableMetaGetAPIResponse) {
	v.Reset()
	poolTaobaoSmartappTableMetaGetAPIResponse.Put(v)
}
