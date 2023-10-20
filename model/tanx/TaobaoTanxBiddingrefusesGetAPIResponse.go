package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxBiddingrefusesGetAPIResponse tanx竞价失败反馈api API返回值
// taobao.tanx.biddingrefuses.get
//
// 竞价失败反馈根据创意id查询API提供
type TaobaoTanxBiddingrefusesGetAPIResponse struct {
	model.CommonResponse
	TaobaoTanxBiddingrefusesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxBiddingrefusesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxBiddingrefusesGetAPIResponseModel).Reset()
}

// TaobaoTanxBiddingrefusesGetAPIResponseModel is tanx竞价失败反馈api 成功返回结果
type TaobaoTanxBiddingrefusesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_biddingrefuses_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回竞价失败对象信息
	BiddingRefuseList []BiddingRefuseDto `json:"bidding_refuse_list,omitempty" xml:"bidding_refuse_list>bidding_refuse_dto,omitempty"`
	// 返回是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTanxBiddingrefusesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BiddingRefuseList = m.BiddingRefuseList[:0]
	m.IsSuccess = false
}

var poolTaobaoTanxBiddingrefusesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxBiddingrefusesGetAPIResponse)
	},
}

// GetTaobaoTanxBiddingrefusesGetAPIResponse 从 sync.Pool 获取 TaobaoTanxBiddingrefusesGetAPIResponse
func GetTaobaoTanxBiddingrefusesGetAPIResponse() *TaobaoTanxBiddingrefusesGetAPIResponse {
	return poolTaobaoTanxBiddingrefusesGetAPIResponse.Get().(*TaobaoTanxBiddingrefusesGetAPIResponse)
}

// ReleaseTaobaoTanxBiddingrefusesGetAPIResponse 将 TaobaoTanxBiddingrefusesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxBiddingrefusesGetAPIResponse(v *TaobaoTanxBiddingrefusesGetAPIResponse) {
	v.Reset()
	poolTaobaoTanxBiddingrefusesGetAPIResponse.Put(v)
}
