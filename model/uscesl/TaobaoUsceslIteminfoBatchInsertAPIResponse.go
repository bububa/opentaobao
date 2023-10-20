package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslIteminfoBatchInsertAPIResponse 按商家批量写入商品接口 API返回值
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
type TaobaoUsceslIteminfoBatchInsertAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslIteminfoBatchInsertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslIteminfoBatchInsertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslIteminfoBatchInsertAPIResponseModel).Reset()
}

// TaobaoUsceslIteminfoBatchInsertAPIResponseModel is 按商家批量写入商品接口 成功返回结果
type TaobaoUsceslIteminfoBatchInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_iteminfo_batch_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslIteminfoBatchInsertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslIteminfoBatchInsertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslIteminfoBatchInsertAPIResponse)
	},
}

// GetTaobaoUsceslIteminfoBatchInsertAPIResponse 从 sync.Pool 获取 TaobaoUsceslIteminfoBatchInsertAPIResponse
func GetTaobaoUsceslIteminfoBatchInsertAPIResponse() *TaobaoUsceslIteminfoBatchInsertAPIResponse {
	return poolTaobaoUsceslIteminfoBatchInsertAPIResponse.Get().(*TaobaoUsceslIteminfoBatchInsertAPIResponse)
}

// ReleaseTaobaoUsceslIteminfoBatchInsertAPIResponse 将 TaobaoUsceslIteminfoBatchInsertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslIteminfoBatchInsertAPIResponse(v *TaobaoUsceslIteminfoBatchInsertAPIResponse) {
	v.Reset()
	poolTaobaoUsceslIteminfoBatchInsertAPIResponse.Put(v)
}
