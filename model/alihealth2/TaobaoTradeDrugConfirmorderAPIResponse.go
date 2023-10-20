package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugConfirmorderAPIResponse 阿里健康020接单 API返回值
// taobao.trade.drug.confirmorder
//
// 阿里健康020接单
type TaobaoTradeDrugConfirmorderAPIResponse struct {
	model.CommonResponse
	TaobaoTradeDrugConfirmorderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeDrugConfirmorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeDrugConfirmorderAPIResponseModel).Reset()
}

// TaobaoTradeDrugConfirmorderAPIResponseModel is 阿里健康020接单 成功返回结果
type TaobaoTradeDrugConfirmorderAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_confirmorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// life的返回值
	Result *LifeResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeDrugConfirmorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTradeDrugConfirmorderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeDrugConfirmorderAPIResponse)
	},
}

// GetTaobaoTradeDrugConfirmorderAPIResponse 从 sync.Pool 获取 TaobaoTradeDrugConfirmorderAPIResponse
func GetTaobaoTradeDrugConfirmorderAPIResponse() *TaobaoTradeDrugConfirmorderAPIResponse {
	return poolTaobaoTradeDrugConfirmorderAPIResponse.Get().(*TaobaoTradeDrugConfirmorderAPIResponse)
}

// ReleaseTaobaoTradeDrugConfirmorderAPIResponse 将 TaobaoTradeDrugConfirmorderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeDrugConfirmorderAPIResponse(v *TaobaoTradeDrugConfirmorderAPIResponse) {
	v.Reset()
	poolTaobaoTradeDrugConfirmorderAPIResponse.Put(v)
}
