package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemImgDeleteAPIResponse 删除商品图片 API返回值
// taobao.item.img.delete
//
// 删除商品图片
type TaobaoItemImgDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoItemImgDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemImgDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemImgDeleteAPIResponseModel).Reset()
}

// TaobaoItemImgDeleteAPIResponseModel is 删除商品图片 成功返回结果
type TaobaoItemImgDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"item_img_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品图片结构
	ItemImg *ItemImg `json:"item_img,omitempty" xml:"item_img,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemImgDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemImg = nil
}

var poolTaobaoItemImgDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemImgDeleteAPIResponse)
	},
}

// GetTaobaoItemImgDeleteAPIResponse 从 sync.Pool 获取 TaobaoItemImgDeleteAPIResponse
func GetTaobaoItemImgDeleteAPIResponse() *TaobaoItemImgDeleteAPIResponse {
	return poolTaobaoItemImgDeleteAPIResponse.Get().(*TaobaoItemImgDeleteAPIResponse)
}

// ReleaseTaobaoItemImgDeleteAPIResponse 将 TaobaoItemImgDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemImgDeleteAPIResponse(v *TaobaoItemImgDeleteAPIResponse) {
	v.Reset()
	poolTaobaoItemImgDeleteAPIResponse.Put(v)
}
