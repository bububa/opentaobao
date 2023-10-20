package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityNameQueryAPIResponse 查询官方的活动名称接口 API返回值
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
type TaobaoSingletreasureActivityNameQueryAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityNameQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityNameQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityNameQueryAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityNameQueryAPIResponseModel is 查询官方的活动名称接口 成功返回结果
type TaobaoSingletreasureActivityNameQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_name_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityNameQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityNameQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityNameQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityNameQueryAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityNameQueryAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityNameQueryAPIResponse
func GetTaobaoSingletreasureActivityNameQueryAPIResponse() *TaobaoSingletreasureActivityNameQueryAPIResponse {
	return poolTaobaoSingletreasureActivityNameQueryAPIResponse.Get().(*TaobaoSingletreasureActivityNameQueryAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityNameQueryAPIResponse 将 TaobaoSingletreasureActivityNameQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityNameQueryAPIResponse(v *TaobaoSingletreasureActivityNameQueryAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityNameQueryAPIResponse.Put(v)
}
