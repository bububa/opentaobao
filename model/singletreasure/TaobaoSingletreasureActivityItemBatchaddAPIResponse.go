package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemBatchaddAPIResponse 批量添加商品接口 API返回值
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
type TaobaoSingletreasureActivityItemBatchaddAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityItemBatchaddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemBatchaddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityItemBatchaddAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityItemBatchaddAPIResponseModel is 批量添加商品接口 成功返回结果
type TaobaoSingletreasureActivityItemBatchaddAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_batchadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityItemBatchaddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemBatchaddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityItemBatchaddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemBatchaddAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityItemBatchaddAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityItemBatchaddAPIResponse
func GetTaobaoSingletreasureActivityItemBatchaddAPIResponse() *TaobaoSingletreasureActivityItemBatchaddAPIResponse {
	return poolTaobaoSingletreasureActivityItemBatchaddAPIResponse.Get().(*TaobaoSingletreasureActivityItemBatchaddAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityItemBatchaddAPIResponse 将 TaobaoSingletreasureActivityItemBatchaddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemBatchaddAPIResponse(v *TaobaoSingletreasureActivityItemBatchaddAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemBatchaddAPIResponse.Put(v)
}
