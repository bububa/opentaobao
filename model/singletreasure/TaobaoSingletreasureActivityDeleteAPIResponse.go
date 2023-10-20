package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityDeleteAPIResponse 删除活动接口 API返回值
// taobao.singletreasure.activity.delete
//
// 删除优惠活动
type TaobaoSingletreasureActivityDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityDeleteAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityDeleteAPIResponseModel is 删除活动接口 成功返回结果
type TaobaoSingletreasureActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 删除是否成功 boolean值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 系统执行是否成功
	ResultStatus bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = 0
	m.Data = false
	m.ResultStatus = false
}

var poolTaobaoSingletreasureActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityDeleteAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityDeleteAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityDeleteAPIResponse
func GetTaobaoSingletreasureActivityDeleteAPIResponse() *TaobaoSingletreasureActivityDeleteAPIResponse {
	return poolTaobaoSingletreasureActivityDeleteAPIResponse.Get().(*TaobaoSingletreasureActivityDeleteAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityDeleteAPIResponse 将 TaobaoSingletreasureActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityDeleteAPIResponse(v *TaobaoSingletreasureActivityDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityDeleteAPIResponse.Put(v)
}
