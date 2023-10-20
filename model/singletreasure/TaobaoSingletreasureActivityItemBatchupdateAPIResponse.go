package singletreasure

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemBatchupdateAPIResponse 批量修改商品接口 API返回值
// taobao.singletreasure.activity.item.batchupdate
//
// 批量修改商品优惠接口
type TaobaoSingletreasureActivityItemBatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaoSingletreasureActivityItemBatchupdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemBatchupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSingletreasureActivityItemBatchupdateAPIResponseModel).Reset()
}

// TaobaoSingletreasureActivityItemBatchupdateAPIResponseModel is 批量修改商品接口 成功返回结果
type TaobaoSingletreasureActivityItemBatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_batchupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSingletreasureActivityItemBatchupdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSingletreasureActivityItemBatchupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSingletreasureActivityItemBatchupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemBatchupdateAPIResponse)
	},
}

// GetTaobaoSingletreasureActivityItemBatchupdateAPIResponse 从 sync.Pool 获取 TaobaoSingletreasureActivityItemBatchupdateAPIResponse
func GetTaobaoSingletreasureActivityItemBatchupdateAPIResponse() *TaobaoSingletreasureActivityItemBatchupdateAPIResponse {
	return poolTaobaoSingletreasureActivityItemBatchupdateAPIResponse.Get().(*TaobaoSingletreasureActivityItemBatchupdateAPIResponse)
}

// ReleaseTaobaoSingletreasureActivityItemBatchupdateAPIResponse 将 TaobaoSingletreasureActivityItemBatchupdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemBatchupdateAPIResponse(v *TaobaoSingletreasureActivityItemBatchupdateAPIResponse) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemBatchupdateAPIResponse.Put(v)
}
