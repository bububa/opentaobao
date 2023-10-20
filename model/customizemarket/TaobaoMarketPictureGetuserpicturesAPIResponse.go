package customizemarket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaomarketpicturegetuserpicturesAPIResponse 读取用户上传图片 API返回值
// taobao.market.picture.getuserpictures
//
// 商家通过用户信息，获取用户上传的
type TaobaomarketpicturegetuserpicturesAPIResponse struct {
	model.CommonResponse
	TaobaomarketpicturegetuserpicturesAPIResponseModel
}

// TaobaomarketpicturegetuserpicturesAPIResponseModel is 读取用户上传图片 成功返回结果
type TaobaomarketpicturegetuserpicturesAPIResponseModel struct {
	XMLName xml.Name `xml:"market_picture_getuserpictures_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
