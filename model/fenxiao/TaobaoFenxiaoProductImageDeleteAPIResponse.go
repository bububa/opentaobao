package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductImageDeleteAPIResponse 产品图片删除 API返回值
// taobao.fenxiao.product.image.delete
//
// 产品图片删除，只删除图片信息，不真正删除图片
type TaobaoFenxiaoProductImageDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductImageDeleteAPIResponseModel
}

// TaobaoFenxiaoProductImageDeleteAPIResponseModel is 产品图片删除 成功返回结果
type TaobaoFenxiaoProductImageDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_image_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
