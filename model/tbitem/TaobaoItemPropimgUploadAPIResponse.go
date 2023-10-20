package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPropimgUploadAPIResponse 添加或修改属性图片 API返回值
// taobao.item.propimg.upload
//
// 添加一张商品属性图片到num_iid指定的商品中 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 &lt;br/&gt;商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 &lt;br/&gt;商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoItemPropimgUploadAPIResponse struct {
	model.CommonResponse
	TaobaoItemPropimgUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemPropimgUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemPropimgUploadAPIResponseModel).Reset()
}

// TaobaoItemPropimgUploadAPIResponseModel is 添加或修改属性图片 成功返回结果
type TaobaoItemPropimgUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"item_propimg_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// PropImg属性图片结构
	PropImg *PropImg `json:"prop_img,omitempty" xml:"prop_img,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemPropimgUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PropImg = nil
}

var poolTaobaoItemPropimgUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemPropimgUploadAPIResponse)
	},
}

// GetTaobaoItemPropimgUploadAPIResponse 从 sync.Pool 获取 TaobaoItemPropimgUploadAPIResponse
func GetTaobaoItemPropimgUploadAPIResponse() *TaobaoItemPropimgUploadAPIResponse {
	return poolTaobaoItemPropimgUploadAPIResponse.Get().(*TaobaoItemPropimgUploadAPIResponse)
}

// ReleaseTaobaoItemPropimgUploadAPIResponse 将 TaobaoItemPropimgUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemPropimgUploadAPIResponse(v *TaobaoItemPropimgUploadAPIResponse) {
	v.Reset()
	poolTaobaoItemPropimgUploadAPIResponse.Put(v)
}
