package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGradeGetAPIResponse 卖家查询等级规则 API返回值
// taobao.crm.grade.get
//
// 卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息
type TaobaoCrmGradeGetAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGradeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGradeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGradeGetAPIResponseModel).Reset()
}

// TaobaoCrmGradeGetAPIResponseModel is 卖家查询等级规则 成功返回结果
type TaobaoCrmGradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 等级信息集合
	GradePromotions []GradePromotion `json:"grade_promotions,omitempty" xml:"grade_promotions>grade_promotion,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGradeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GradePromotions = m.GradePromotions[:0]
}

var poolTaobaoCrmGradeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGradeGetAPIResponse)
	},
}

// GetTaobaoCrmGradeGetAPIResponse 从 sync.Pool 获取 TaobaoCrmGradeGetAPIResponse
func GetTaobaoCrmGradeGetAPIResponse() *TaobaoCrmGradeGetAPIResponse {
	return poolTaobaoCrmGradeGetAPIResponse.Get().(*TaobaoCrmGradeGetAPIResponse)
}

// ReleaseTaobaoCrmGradeGetAPIResponse 将 TaobaoCrmGradeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGradeGetAPIResponse(v *TaobaoCrmGradeGetAPIResponse) {
	v.Reset()
	poolTaobaoCrmGradeGetAPIResponse.Put(v)
}
