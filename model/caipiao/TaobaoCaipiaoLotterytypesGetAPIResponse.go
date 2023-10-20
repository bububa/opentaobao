package caipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoLotterytypesGetAPIResponse 获取可用的彩种列表 API返回值
// taobao.caipiao.lotterytypes.get
//
// 获取彩票系统支持的可用于赠送的彩种列表
type TaobaoCaipiaoLotterytypesGetAPIResponse struct {
	model.CommonResponse
	TaobaoCaipiaoLotterytypesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCaipiaoLotterytypesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCaipiaoLotterytypesGetAPIResponseModel).Reset()
}

// TaobaoCaipiaoLotterytypesGetAPIResponseModel is 获取可用的彩种列表 成功返回结果
type TaobaoCaipiaoLotterytypesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"caipiao_lotterytypes_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 彩种的数据结构
	Results []LotteryType `json:"results,omitempty" xml:"results>lottery_type,omitempty"`
	// 彩种个数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCaipiaoLotterytypesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.TotalResults = 0
}

var poolTaobaoCaipiaoLotterytypesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCaipiaoLotterytypesGetAPIResponse)
	},
}

// GetTaobaoCaipiaoLotterytypesGetAPIResponse 从 sync.Pool 获取 TaobaoCaipiaoLotterytypesGetAPIResponse
func GetTaobaoCaipiaoLotterytypesGetAPIResponse() *TaobaoCaipiaoLotterytypesGetAPIResponse {
	return poolTaobaoCaipiaoLotterytypesGetAPIResponse.Get().(*TaobaoCaipiaoLotterytypesGetAPIResponse)
}

// ReleaseTaobaoCaipiaoLotterytypesGetAPIResponse 将 TaobaoCaipiaoLotterytypesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCaipiaoLotterytypesGetAPIResponse(v *TaobaoCaipiaoLotterytypesGetAPIResponse) {
	v.Reset()
	poolTaobaoCaipiaoLotterytypesGetAPIResponse.Put(v)
}
