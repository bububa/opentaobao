package customizemarket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMarketPictureGetuserpicturesAPIResponse 读取用户上传图片 API返回值
// taobao.market.picture.getuserpictures
//
// 商家通过用户信息，获取用户上传的
type TaobaoMarketPictureGetuserpicturesAPIResponse struct {
	model.CommonResponse
	TaobaoMarketPictureGetuserpicturesAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMarketPictureGetuserpicturesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMarketPictureGetuserpicturesAPIResponseModel).Reset()
}

// TaobaoMarketPictureGetuserpicturesAPIResponseModel is 读取用户上传图片 成功返回结果
type TaobaoMarketPictureGetuserpicturesAPIResponseModel struct {
	XMLName xml.Name `xml:"market_picture_getuserpictures_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMarketPictureGetuserpicturesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMarketPictureGetuserpicturesAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMarketPictureGetuserpicturesAPIResponse)
	},
}

// GetTaobaoMarketPictureGetuserpicturesAPIResponse 从 sync.Pool 获取 TaobaoMarketPictureGetuserpicturesAPIResponse
func GetTaobaoMarketPictureGetuserpicturesAPIResponse() *TaobaoMarketPictureGetuserpicturesAPIResponse {
	return poolTaobaoMarketPictureGetuserpicturesAPIResponse.Get().(*TaobaoMarketPictureGetuserpicturesAPIResponse)
}

// ReleaseTaobaoMarketPictureGetuserpicturesAPIResponse 将 TaobaoMarketPictureGetuserpicturesAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMarketPictureGetuserpicturesAPIResponse(v *TaobaoMarketPictureGetuserpicturesAPIResponse) {
	v.Reset()
	poolTaobaoMarketPictureGetuserpicturesAPIResponse.Put(v)
}
