package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemImgDeleteAPIResponse
删除商品图片 API返回值
taobao.item.img.delete

删除商品图片 */
type TaobaoItemImgDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoItemImgDeleteAPIResponseModel
}

// TaobaoItemImgDeleteAPIResponseModel is 删除商品图片 成功返回结果
type TaobaoItemImgDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"item_img_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品图片结构
	ItemImg *ItemImg `json:"item_img,omitempty" xml:"item_img,omitempty"`
}
