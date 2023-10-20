package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateListingTmallAPIResponse taobao.item.update.listing天猫分流 API返回值
// taobao.item.update.listing.tmall
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingTmallAPIResponse struct {
	model.CommonResponse
	TaobaoItemUpdateListingTmallAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemUpdateListingTmallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemUpdateListingTmallAPIResponseModel).Reset()
}

// TaobaoItemUpdateListingTmallAPIResponseModel is taobao.item.update.listing天猫分流 成功返回结果
type TaobaoItemUpdateListingTmallAPIResponseModel struct {
	XMLName xml.Name `xml:"item_update_listing_tmall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上架后返回的商品信息：返回的结果就是:num_iid和modified
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemUpdateListingTmallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Item = nil
}

var poolTaobaoItemUpdateListingTmallAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemUpdateListingTmallAPIResponse)
	},
}

// GetTaobaoItemUpdateListingTmallAPIResponse 从 sync.Pool 获取 TaobaoItemUpdateListingTmallAPIResponse
func GetTaobaoItemUpdateListingTmallAPIResponse() *TaobaoItemUpdateListingTmallAPIResponse {
	return poolTaobaoItemUpdateListingTmallAPIResponse.Get().(*TaobaoItemUpdateListingTmallAPIResponse)
}

// ReleaseTaobaoItemUpdateListingTmallAPIResponse 将 TaobaoItemUpdateListingTmallAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemUpdateListingTmallAPIResponse(v *TaobaoItemUpdateListingTmallAPIResponse) {
	v.Reset()
	poolTaobaoItemUpdateListingTmallAPIResponse.Put(v)
}
