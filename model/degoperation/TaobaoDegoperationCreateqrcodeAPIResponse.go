package degoperation

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationCreateqrcodeAPIResponse 中奖生成二维码 API返回值
// taobao.degoperation.createqrcode
//
// 用户中奖后，生成二维码图片链接
type TaobaoDegoperationCreateqrcodeAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationCreateqrcodeAPIResponseModel
}

// TaobaoDegoperationCreateqrcodeAPIResponseModel is 中奖生成二维码 成功返回结果
type TaobaoDegoperationCreateqrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_createqrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 二维码链接
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// realSubCode
	RealSubCode string `json:"real_sub_code,omitempty" xml:"real_sub_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
