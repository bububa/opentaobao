package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateRelationshipwithrpGetAPIResponse 根据gid查询卖家下所有的rpId API返回值
// taobao.xhotel.rate.relationshipwithrp.get
//
// 根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
type TaobaoXhotelRateRelationshipwithrpGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateRelationshipwithrpGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateRelationshipwithrpGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateRelationshipwithrpGetAPIResponseModel).Reset()
}

// TaobaoXhotelRateRelationshipwithrpGetAPIResponseModel is 根据gid查询卖家下所有的rpId 成功返回结果
type TaobaoXhotelRateRelationshipwithrpGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_relationshipwithrp_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 所查询出的结果，是一个字符串数组
	RpIds []string `json:"rp_ids,omitempty" xml:"rp_ids>string,omitempty"`
	// 根据条件所查询的所有结果的总数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateRelationshipwithrpGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RpIds = m.RpIds[:0]
	m.TotalResults = 0
}

var poolTaobaoXhotelRateRelationshipwithrpGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateRelationshipwithrpGetAPIResponse)
	},
}

// GetTaobaoXhotelRateRelationshipwithrpGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateRelationshipwithrpGetAPIResponse
func GetTaobaoXhotelRateRelationshipwithrpGetAPIResponse() *TaobaoXhotelRateRelationshipwithrpGetAPIResponse {
	return poolTaobaoXhotelRateRelationshipwithrpGetAPIResponse.Get().(*TaobaoXhotelRateRelationshipwithrpGetAPIResponse)
}

// ReleaseTaobaoXhotelRateRelationshipwithrpGetAPIResponse 将 TaobaoXhotelRateRelationshipwithrpGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateRelationshipwithrpGetAPIResponse(v *TaobaoXhotelRateRelationshipwithrpGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateRelationshipwithrpGetAPIResponse.Put(v)
}
