package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIResponse 根据用户数字id和身份证号校验该用户是否已实名认证成功 API返回值
// taobao.paimai.auctioncat.nft.checknftuseridentify
//
// 根据用户数字id和身份证号校验该用户是否已实名认证成功
type TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIResponse struct {
	model.CommonResponse
	TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIResponseModel
}

// TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIResponseModel is 根据用户数字id和身份证号校验该用户是否已实名认证成功 成功返回结果
type TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_auctioncat_nft_checknftuseridentify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 验证是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
