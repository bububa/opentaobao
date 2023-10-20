package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonMallDetailGetAPIResponse 查询商圈详细信息 API返回值
// taobao.koubei.mall.common.mall.detail.get
//
// 查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
type TaobaoKoubeiMallCommonMallDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonMallDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonMallDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonMallDetailGetAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonMallDetailGetAPIResponseModel is 查询商圈详细信息 成功返回结果
type TaobaoKoubeiMallCommonMallDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_mall_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaoKoubeiMallCommonMallDetailGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonMallDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonMallDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonMallDetailGetAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonMallDetailGetAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonMallDetailGetAPIResponse
func GetTaobaoKoubeiMallCommonMallDetailGetAPIResponse() *TaobaoKoubeiMallCommonMallDetailGetAPIResponse {
	return poolTaobaoKoubeiMallCommonMallDetailGetAPIResponse.Get().(*TaobaoKoubeiMallCommonMallDetailGetAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonMallDetailGetAPIResponse 将 TaobaoKoubeiMallCommonMallDetailGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonMallDetailGetAPIResponse(v *TaobaoKoubeiMallCommonMallDetailGetAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonMallDetailGetAPIResponse.Put(v)
}
