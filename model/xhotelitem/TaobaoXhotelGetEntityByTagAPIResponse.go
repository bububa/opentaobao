package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelGetEntityByTagAPIResponse 根据标签查询实体 API返回值
// taobao.xhotel.get.entity.by.tag
//
// 根据标签查询实体
type TaobaoXhotelGetEntityByTagAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelGetEntityByTagAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelGetEntityByTagAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelGetEntityByTagAPIResponseModel).Reset()
}

// TaobaoXhotelGetEntityByTagAPIResponseModel is 根据标签查询实体 成功返回结果
type TaobaoXhotelGetEntityByTagAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_get_entity_by_tag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	TagQueryResult *TagQueryResult `json:"tag_query_result,omitempty" xml:"tag_query_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelGetEntityByTagAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TagQueryResult = nil
}

var poolTaobaoXhotelGetEntityByTagAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelGetEntityByTagAPIResponse)
	},
}

// GetTaobaoXhotelGetEntityByTagAPIResponse 从 sync.Pool 获取 TaobaoXhotelGetEntityByTagAPIResponse
func GetTaobaoXhotelGetEntityByTagAPIResponse() *TaobaoXhotelGetEntityByTagAPIResponse {
	return poolTaobaoXhotelGetEntityByTagAPIResponse.Get().(*TaobaoXhotelGetEntityByTagAPIResponse)
}

// ReleaseTaobaoXhotelGetEntityByTagAPIResponse 将 TaobaoXhotelGetEntityByTagAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelGetEntityByTagAPIResponse(v *TaobaoXhotelGetEntityByTagAPIResponse) {
	v.Reset()
	poolTaobaoXhotelGetEntityByTagAPIResponse.Put(v)
}
