package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemQueryAPIResponse 查询活动下的优惠信息 API返回值
// taobao.singletreasure.activity.item.query
//
// 分页查询活动下的商品优惠信息
type TaobaoSingletreasureActivityItemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityItemQueryAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityItemQueryAPIResponseModel is 查询活动下的优惠信息 成功返回结果
type TaobaoSingletreasureActivityItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemQueryAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityItemQueryAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityItemQueryAPIResponse
func GetTaobaoSingletreasureActivityItemQueryAPIResponse() *TaobaoSingletreasureActivityItemQueryAPIResponse {
	return poolTaobaoSingletreasureActivityItemQueryAPIResponse.Get().(*TaobaoSingletreasureActivityItemQueryAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityItemQueryAPIResponse 将 TaobaoSingletreasureActivityItemQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemQueryAPIResponse(v *TaobaoSingletreasureActivityItemQueryAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemQueryAPIResponse.Put(v)
}
