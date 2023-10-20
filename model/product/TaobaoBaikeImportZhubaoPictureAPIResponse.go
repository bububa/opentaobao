package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaikeimportzhubaopictureAPIResponse 百科图片数据导入 API返回值
// taobao.baike.import.zhubao.picture
//
// 用于接入外部--图片--录入到商品百科中
type TaobaobaikeimportzhubaopictureAPIResponse struct {
	model.CommonResponse
	TaobaobaikeimportzhubaopictureAPIResponseModel
}

// TaobaobaikeimportzhubaopictureAPIResponseModel is 百科图片数据导入 成功返回结果
type TaobaobaikeimportzhubaopictureAPIResponseModel struct {
	XMLName xml.Name `xml:"baike_import_zhubao_picture_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaobaikeimportzhubaopictureResult `json:"result,omitempty" xml:"result,omitempty"`
}
