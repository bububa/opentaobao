package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityUpdateAPIResponse 修改活动接口 API返回值
// taobao.singletreasure.activity.update
//
// 修改活动接口
type TaobaoSingletreasureActivityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityUpdateAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityUpdateAPIResponseModel is 修改活动接口 成功返回结果
type TaobaoSingletreasureActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityUpdateAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityUpdateAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityUpdateAPIResponse
func GetTaobaoSingletreasureActivityUpdateAPIResponse() *TaobaoSingletreasureActivityUpdateAPIResponse {
	return poolTaobaoSingletreasureActivityUpdateAPIResponse.Get().(*TaobaoSingletreasureActivityUpdateAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityUpdateAPIResponse 将 TaobaoSingletreasureActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityUpdateAPIResponse(v *TaobaoSingletreasureActivityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityUpdateAPIResponse.Put(v)
}
