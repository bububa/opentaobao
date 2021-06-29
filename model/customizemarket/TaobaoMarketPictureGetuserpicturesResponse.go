package customizemarket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
读取用户上传图片 APIResponse
taobao.market.picture.getuserpictures

商家通过用户信息，获取用户上传的
*/
type TaobaoMarketPictureGetuserpicturesAPIResponse struct {
    model.CommonResponse
    TaobaoMarketPictureGetuserpicturesResponse
}

type TaobaoMarketPictureGetuserpicturesResponse struct {
    XMLName xml.Name `xml:"market_picture_getuserpictures_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
