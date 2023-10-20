package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappSmartformDataWriteAPIResponse 智能表单外部更新数据 API返回值
// taobao.smartapp.smartform.data.write
//
// 智能表单外部更新数据
type TaobaoSmartappSmartformDataWriteAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappSmartformDataWriteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartappSmartformDataWriteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartappSmartformDataWriteAPIResponseModel).Reset()
}

// TaobaoSmartappSmartformDataWriteAPIResponseModel is 智能表单外部更新数据 成功返回结果
type TaobaoSmartappSmartformDataWriteAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_smartform_data_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作数据ID
	RecordId string `json:"record_id,omitempty" xml:"record_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSmartappSmartformDataWriteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RecordId = ""
}

var poolTaobaoSmartappSmartformDataWriteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartappSmartformDataWriteAPIResponse)
	},
}

// GetTaobaoSmartappSmartformDataWriteAPIResponse 从 sync.Pool 获取 TaobaoSmartappSmartformDataWriteAPIResponse
func GetTaobaoSmartappSmartformDataWriteAPIResponse() *TaobaoSmartappSmartformDataWriteAPIResponse {
	return poolTaobaoSmartappSmartformDataWriteAPIResponse.Get().(*TaobaoSmartappSmartformDataWriteAPIResponse)
}

// ReleaseTaobaoSmartappSmartformDataWriteAPIResponse 将 TaobaoSmartappSmartformDataWriteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartappSmartformDataWriteAPIResponse(v *TaobaoSmartappSmartformDataWriteAPIResponse) {
	v.Reset()
	poolTaobaoSmartappSmartformDataWriteAPIResponse.Put(v)
}
