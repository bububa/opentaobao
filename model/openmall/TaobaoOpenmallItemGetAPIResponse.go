package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemGetAPIResponse 获取商品详情物料 API返回值
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
type TaobaoOpenmallItemGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallItemGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallItemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallItemGetAPIResponseModel).Reset()
}

// TaobaoOpenmallItemGetAPIResponseModel is 获取商品详情物料 成功返回结果
type TaobaoOpenmallItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoOpenmallItemGetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallItemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallItemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallItemGetAPIResponse)
	},
}

// GetTaobaoOpenmallItemGetAPIResponse 从 sync.Pool 获取 TaobaoOpenmallItemGetAPIResponse
func GetTaobaoOpenmallItemGetAPIResponse() *TaobaoOpenmallItemGetAPIResponse {
	return poolTaobaoOpenmallItemGetAPIResponse.Get().(*TaobaoOpenmallItemGetAPIResponse)
}

// ReleaseTaobaoOpenmallItemGetAPIResponse 将 TaobaoOpenmallItemGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallItemGetAPIResponse(v *TaobaoOpenmallItemGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallItemGetAPIResponse.Put(v)
}
