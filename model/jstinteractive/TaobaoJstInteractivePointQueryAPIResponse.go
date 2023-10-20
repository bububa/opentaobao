package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointQueryAPIResponse 互动积分查询接口 API返回值
// taobao.jst.interactive.point.query
//
// 查询用户的互动积分
type TaobaoJstInteractivePointQueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractivePointQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractivePointQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractivePointQueryAPIResponseModel).Reset()
}

// TaobaoJstInteractivePointQueryAPIResponseModel is 互动积分查询接口 成功返回结果
type TaobaoJstInteractivePointQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_point_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Data *InteractivePointQueryResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractivePointQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolTaobaoJstInteractivePointQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractivePointQueryAPIResponse)
	},
}

// GetTaobaoJstInteractivePointQueryAPIResponse 从 sync.Pool 获取 TaobaoJstInteractivePointQueryAPIResponse
func GetTaobaoJstInteractivePointQueryAPIResponse() *TaobaoJstInteractivePointQueryAPIResponse {
	return poolTaobaoJstInteractivePointQueryAPIResponse.Get().(*TaobaoJstInteractivePointQueryAPIResponse)
}

// ReleaseTaobaoJstInteractivePointQueryAPIResponse 将 TaobaoJstInteractivePointQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractivePointQueryAPIResponse(v *TaobaoJstInteractivePointQueryAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractivePointQueryAPIResponse.Put(v)
}
