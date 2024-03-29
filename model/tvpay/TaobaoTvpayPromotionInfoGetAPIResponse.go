package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayPromotionInfoGetAPIResponse tv支付查询消费抽奖配置 API返回值
// taobao.tvpay.promotion.info.get
//
// 查询消费抽奖配置
type TaobaoTvpayPromotionInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayPromotionInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayPromotionInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayPromotionInfoGetAPIResponseModel).Reset()
}

// TaobaoTvpayPromotionInfoGetAPIResponseModel is tv支付查询消费抽奖配置 成功返回结果
type TaobaoTvpayPromotionInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_promotion_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayPromotionInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayPromotionInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayPromotionInfoGetAPIResponse)
	},
}

// GetTaobaoTvpayPromotionInfoGetAPIResponse 从 sync.Pool 获取 TaobaoTvpayPromotionInfoGetAPIResponse
func GetTaobaoTvpayPromotionInfoGetAPIResponse() *TaobaoTvpayPromotionInfoGetAPIResponse {
	return poolTaobaoTvpayPromotionInfoGetAPIResponse.Get().(*TaobaoTvpayPromotionInfoGetAPIResponse)
}

// ReleaseTaobaoTvpayPromotionInfoGetAPIResponse 将 TaobaoTvpayPromotionInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayPromotionInfoGetAPIResponse(v *TaobaoTvpayPromotionInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoTvpayPromotionInfoGetAPIResponse.Put(v)
}
