package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductPropimgUploadAPIResponse 上传单张产品属性图片，如果需要传多张，可调多次 API返回值
// taobao.product.propimg.upload
//
// 传入产品ID &lt;br/&gt;传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid,&lt;br/&gt;再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; &lt;br/&gt;传入图片内容 &lt;br/&gt;注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次
type TaobaoProductPropimgUploadAPIResponse struct {
	model.CommonResponse
	TaobaoProductPropimgUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoProductPropimgUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoProductPropimgUploadAPIResponseModel).Reset()
}

// TaobaoProductPropimgUploadAPIResponseModel is 上传单张产品属性图片，如果需要传多张，可调多次 成功返回结果
type TaobaoProductPropimgUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"product_propimg_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 支持返回产品属性图片中的：url,id,created,modified
	ProductPropImg *ProductPropImg `json:"product_prop_img,omitempty" xml:"product_prop_img,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoProductPropimgUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductPropImg = nil
}

var poolTaobaoProductPropimgUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoProductPropimgUploadAPIResponse)
	},
}

// GetTaobaoProductPropimgUploadAPIResponse 从 sync.Pool 获取 TaobaoProductPropimgUploadAPIResponse
func GetTaobaoProductPropimgUploadAPIResponse() *TaobaoProductPropimgUploadAPIResponse {
	return poolTaobaoProductPropimgUploadAPIResponse.Get().(*TaobaoProductPropimgUploadAPIResponse)
}

// ReleaseTaobaoProductPropimgUploadAPIResponse 将 TaobaoProductPropimgUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoProductPropimgUploadAPIResponse(v *TaobaoProductPropimgUploadAPIResponse) {
	v.Reset()
	poolTaobaoProductPropimgUploadAPIResponse.Put(v)
}
