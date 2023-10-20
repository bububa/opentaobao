package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbreviewaddAPIResponse 对外开放评论接口 API返回值
// taobao.xhotel.bnbreview.add
//
// 对外开放评论接口
type TaobaoxhotelbnbreviewaddAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbreviewaddAPIResponseModel
}

// TaobaoxhotelbnbreviewaddAPIResponseModel is 对外开放评论接口 成功返回结果
type TaobaoxhotelbnbreviewaddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbreview_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用返回结果
	Result *BnbResult `json:"result,omitempty" xml:"result,omitempty"`
}
