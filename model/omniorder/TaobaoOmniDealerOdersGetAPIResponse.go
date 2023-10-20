package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniDealerOdersGetAPIResponse 获取单笔全渠道经销商订单的详细信息 API返回值
// taobao.omni.dealer.oders.get
//
// 全渠道经销商获取单笔订单的详细信息
type TaobaoOmniDealerOdersGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniDealerOdersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniDealerOdersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniDealerOdersGetAPIResponseModel).Reset()
}

// TaobaoOmniDealerOdersGetAPIResponseModel is 获取单笔全渠道经销商订单的详细信息 成功返回结果
type TaobaoOmniDealerOdersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_dealer_oders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单
	Data *TaobaoOmniDealerOdersGetData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniDealerOdersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoOmniDealerOdersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniDealerOdersGetAPIResponse)
	},
}

// GetTaobaoOmniDealerOdersGetAPIResponse 从 sync.Pool 获取 TaobaoOmniDealerOdersGetAPIResponse
func GetTaobaoOmniDealerOdersGetAPIResponse() *TaobaoOmniDealerOdersGetAPIResponse {
	return poolTaobaoOmniDealerOdersGetAPIResponse.Get().(*TaobaoOmniDealerOdersGetAPIResponse)
}

// ReleaseTaobaoOmniDealerOdersGetAPIResponse 将 TaobaoOmniDealerOdersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniDealerOdersGetAPIResponse(v *TaobaoOmniDealerOdersGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniDealerOdersGetAPIResponse.Put(v)
}
