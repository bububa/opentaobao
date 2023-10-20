package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemDeleteAPIResponse 删除单品优惠接口 API返回值
// taobao.singletreasure.activity.item.delete
//
// 删除单品优惠接口
type TaobaoSingletreasureActivityItemDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityItemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityItemDeleteAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityItemDeleteAPIResponseModel is 删除单品优惠接口 成功返回结果
type TaobaoSingletreasureActivityItemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityItemDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityItemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemDeleteAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityItemDeleteAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityItemDeleteAPIResponse
func GetTaobaoSingletreasureActivityItemDeleteAPIResponse() *TaobaoSingletreasureActivityItemDeleteAPIResponse {
	return poolTaobaoSingletreasureActivityItemDeleteAPIResponse.Get().(*TaobaoSingletreasureActivityItemDeleteAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityItemDeleteAPIResponse 将 TaobaoSingletreasureActivityItemDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemDeleteAPIResponse(v *TaobaoSingletreasureActivityItemDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemDeleteAPIResponse.Put(v)
}
