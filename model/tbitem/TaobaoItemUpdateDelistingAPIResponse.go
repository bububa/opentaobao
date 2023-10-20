package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateDelistingAPIResponse 商品下架 API返回值
// taobao.item.update.delisting
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingAPIResponse struct {
	model.CommonResponse
	TaobaoItemUpdateDelistingAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemUpdateDelistingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemUpdateDelistingAPIResponseModel).Reset()
}

// TaobaoItemUpdateDelistingAPIResponseModel is 商品下架 成功返回结果
type TaobaoItemUpdateDelistingAPIResponseModel struct {
	XMLName xml.Name `xml:"item_update_delisting_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回商品更新信息：返回的结果是:num_iid和modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemUpdateDelistingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Item = nil
}

var poolTaobaoItemUpdateDelistingAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemUpdateDelistingAPIResponse)
	},
}

// GetTaobaoItemUpdateDelistingAPIResponse 从 sync.Pool 获取 TaobaoItemUpdateDelistingAPIResponse
func GetTaobaoItemUpdateDelistingAPIResponse() *TaobaoItemUpdateDelistingAPIResponse {
	return poolTaobaoItemUpdateDelistingAPIResponse.Get().(*TaobaoItemUpdateDelistingAPIResponse)
}

// ReleaseTaobaoItemUpdateDelistingAPIResponse 将 TaobaoItemUpdateDelistingAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemUpdateDelistingAPIResponse(v *TaobaoItemUpdateDelistingAPIResponse) {
	v.Reset()
	poolTaobaoItemUpdateDelistingAPIResponse.Put(v)
}
