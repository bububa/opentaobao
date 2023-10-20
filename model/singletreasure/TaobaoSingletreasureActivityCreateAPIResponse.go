package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityCreateAPIResponse 活动创建接口 API返回值
// taobao.singletreasure.activity.create
//
// 创建优惠活动
type TaobaoSingletreasureActivityCreateAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityCreateAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityCreateAPIResponseModel is 活动创建接口 成功返回结果
type TaobaoSingletreasureActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityCreateAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityCreateAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityCreateAPIResponse
func GetTaobaoSingletreasureActivityCreateAPIResponse() *TaobaoSingletreasureActivityCreateAPIResponse {
	return poolTaobaoSingletreasureActivityCreateAPIResponse.Get().(*TaobaoSingletreasureActivityCreateAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityCreateAPIResponse 将 TaobaoSingletreasureActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityCreateAPIResponse(v *TaobaoSingletreasureActivityCreateAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityCreateAPIResponse.Put(v)
}
