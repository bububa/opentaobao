package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemUpdateAPIResponse 更新单品优惠接口 API返回值
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
type TaobaoSingletreasureActivityItemUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityItemUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityItemUpdateAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityItemUpdateAPIResponseModel is 更新单品优惠接口 成功返回结果
type TaobaoSingletreasureActivityItemUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityItemUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityItemUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemUpdateAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityItemUpdateAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityItemUpdateAPIResponse
func GetTaobaoSingletreasureActivityItemUpdateAPIResponse() *TaobaoSingletreasureActivityItemUpdateAPIResponse {
	return poolTaobaoSingletreasureActivityItemUpdateAPIResponse.Get().(*TaobaoSingletreasureActivityItemUpdateAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityItemUpdateAPIResponse 将 TaobaoSingletreasureActivityItemUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemUpdateAPIResponse(v *TaobaoSingletreasureActivityItemUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemUpdateAPIResponse.Put(v)
}
