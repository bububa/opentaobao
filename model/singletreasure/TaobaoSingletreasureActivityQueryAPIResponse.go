package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityQueryAPIResponse 查询活动列表接口 API返回值
// taobao.singletreasure.activity.query
//
// 查询活动列表接口
type TaobaoSingletreasureActivityQueryAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityQueryAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityQueryAPIResponseModel is 查询活动列表接口 成功返回结果
type TaobaoSingletreasureActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoSingletreasureActivityQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityQueryAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityQueryAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityQueryAPIResponse
func GetTaobaoSingletreasureActivityQueryAPIResponse() *TaobaoSingletreasureActivityQueryAPIResponse {
	return poolTaobaoSingletreasureActivityQueryAPIResponse.Get().(*TaobaoSingletreasureActivityQueryAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityQueryAPIResponse 将 TaobaoSingletreasureActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityQueryAPIResponse(v *TaobaoSingletreasureActivityQueryAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityQueryAPIResponse.Put(v)
}
