package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFinanceDetailGetAPIResponse 查询汽车金融订单信息 API返回值
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
type TmallCarFinanceDetailGetAPIResponse struct {
	model.CommonResponse
	TmallCarFinanceDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarFinanceDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarFinanceDetailGetAPIResponseModel).Reset()
}

// TmallCarFinanceDetailGetAPIResponseModel is 查询汽车金融订单信息 成功返回结果
type TmallCarFinanceDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_finance_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 查询数据对象
	Data *FinanceDetailInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarFinanceDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Data = nil
	m.ResultSuccess = false
}

var poolTmallCarFinanceDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarFinanceDetailGetAPIResponse)
	},
}

// GetTmallCarFinanceDetailGetAPIResponse 从 sync.Pool 获取 TmallCarFinanceDetailGetAPIResponse
func GetTmallCarFinanceDetailGetAPIResponse() *TmallCarFinanceDetailGetAPIResponse {
	return poolTmallCarFinanceDetailGetAPIResponse.Get().(*TmallCarFinanceDetailGetAPIResponse)
}

// ReleaseTmallCarFinanceDetailGetAPIResponse 将 TmallCarFinanceDetailGetAPIResponse 保存到 sync.Pool
func ReleaseTmallCarFinanceDetailGetAPIResponse(v *TmallCarFinanceDetailGetAPIResponse) {
	v.Reset()
	poolTmallCarFinanceDetailGetAPIResponse.Put(v)
}
