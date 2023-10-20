package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniDealerOdersListAPIResponse 全渠道经销商订单列表 API返回值
// taobao.omni.dealer.oders.list
//
// 全渠道经销商订单列表查询
type TaobaoOmniDealerOdersListAPIResponse struct {
	model.CommonResponse
	TaobaoOmniDealerOdersListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniDealerOdersListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniDealerOdersListAPIResponseModel).Reset()
}

// TaobaoOmniDealerOdersListAPIResponseModel is 全渠道经销商订单列表 成功返回结果
type TaobaoOmniDealerOdersListAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_dealer_oders_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniDealerOdersListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniDealerOdersListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniDealerOdersListAPIResponse)
	},
}

// GetTaobaoOmniDealerOdersListAPIResponse 从 sync.Pool 获取 TaobaoOmniDealerOdersListAPIResponse
func GetTaobaoOmniDealerOdersListAPIResponse() *TaobaoOmniDealerOdersListAPIResponse {
	return poolTaobaoOmniDealerOdersListAPIResponse.Get().(*TaobaoOmniDealerOdersListAPIResponse)
}

// ReleaseTaobaoOmniDealerOdersListAPIResponse 将 TaobaoOmniDealerOdersListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniDealerOdersListAPIResponse(v *TaobaoOmniDealerOdersListAPIResponse) {
	v.Reset()
	poolTaobaoOmniDealerOdersListAPIResponse.Put(v)
}
