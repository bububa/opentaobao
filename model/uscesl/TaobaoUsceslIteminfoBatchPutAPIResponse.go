package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslIteminfoBatchPutAPIResponse 批量写入商品信息接口 API返回值
// taobao.uscesl.iteminfo.batch.put
//
// 电子架签批量写入商品数据，用于电子价签展示
type TaobaoUsceslIteminfoBatchPutAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslIteminfoBatchPutAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslIteminfoBatchPutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslIteminfoBatchPutAPIResponseModel).Reset()
}

// TaobaoUsceslIteminfoBatchPutAPIResponseModel is 批量写入商品信息接口 成功返回结果
type TaobaoUsceslIteminfoBatchPutAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_iteminfo_batch_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslIteminfoBatchPutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslIteminfoBatchPutAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslIteminfoBatchPutAPIResponse)
	},
}

// GetTaobaoUsceslIteminfoBatchPutAPIResponse 从 sync.Pool 获取 TaobaoUsceslIteminfoBatchPutAPIResponse
func GetTaobaoUsceslIteminfoBatchPutAPIResponse() *TaobaoUsceslIteminfoBatchPutAPIResponse {
	return poolTaobaoUsceslIteminfoBatchPutAPIResponse.Get().(*TaobaoUsceslIteminfoBatchPutAPIResponse)
}

// ReleaseTaobaoUsceslIteminfoBatchPutAPIResponse 将 TaobaoUsceslIteminfoBatchPutAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslIteminfoBatchPutAPIResponse(v *TaobaoUsceslIteminfoBatchPutAPIResponse) {
	v.Reset()
	poolTaobaoUsceslIteminfoBatchPutAPIResponse.Put(v)
}
