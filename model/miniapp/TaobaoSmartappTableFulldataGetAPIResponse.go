package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableFulldataGetAPIResponse 智能应用工作表地址查询 API返回值
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
type TaobaoSmartappTableFulldataGetAPIResponse struct {
	model.CommonResponse
	TaobaoSmartappTableFulldataGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSmartappTableFulldataGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSmartappTableFulldataGetAPIResponseModel).Reset()
}

// TaobaoSmartappTableFulldataGetAPIResponseModel is 智能应用工作表地址查询 成功返回结果
type TaobaoSmartappTableFulldataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"smartapp_table_fulldata_get_response"`
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
func (m *TaobaoSmartappTableFulldataGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = nil
	m.BusinessSuccess = false
}

var poolTaobaoSmartappTableFulldataGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSmartappTableFulldataGetAPIResponse)
	},
}

// GetTaobaoSmartappTableFulldataGetAPIResponse 从 sync.Pool 获取 TaobaoSmartappTableFulldataGetAPIResponse
func GetTaobaoSmartappTableFulldataGetAPIResponse() *TaobaoSmartappTableFulldataGetAPIResponse {
	return poolTaobaoSmartappTableFulldataGetAPIResponse.Get().(*TaobaoSmartappTableFulldataGetAPIResponse)
}

// ReleaseTaobaoSmartappTableFulldataGetAPIResponse 将 TaobaoSmartappTableFulldataGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSmartappTableFulldataGetAPIResponse(v *TaobaoSmartappTableFulldataGetAPIResponse) {
	v.Reset()
	poolTaobaoSmartappTableFulldataGetAPIResponse.Put(v)
}
