package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbreviewAddAPIResponse 对外开放评论接口 API返回值
// taobao.xhotel.bnbreview.add
//
// 对外开放评论接口
type TaobaoXhotelBnbreviewAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbreviewAddAPIResponseModel
}

// TaobaoXhotelBnbreviewAddAPIResponseModel is 对外开放评论接口 成功返回结果
type TaobaoXhotelBnbreviewAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbreview_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用返回结果
	Result *BnbResult `json:"result,omitempty" xml:"result,omitempty"`
}
