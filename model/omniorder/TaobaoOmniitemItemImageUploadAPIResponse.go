package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemImageUploadAPIResponse 全渠道商品上传图片 API返回值
// taobao.omniitem.item.image.upload
//
// 全渠道商品上传图片
type TaobaoOmniitemItemImageUploadAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemItemImageUploadAPIResponseModel
}

// TaobaoOmniitemItemImageUploadAPIResponseModel is 全渠道商品上传图片 成功返回结果
type TaobaoOmniitemItemImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniitemItemImageUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
