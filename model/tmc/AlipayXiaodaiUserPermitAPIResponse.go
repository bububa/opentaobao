package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlipayxiaodaiuserpermitAPIResponse 阿里金融用户授权 API返回值
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
type AlipayxiaodaiuserpermitAPIResponse struct {
	model.CommonResponse
	AlipayxiaodaiuserpermitAPIResponseModel
}

// AlipayxiaodaiuserpermitAPIResponseModel is 阿里金融用户授权 成功返回结果
type AlipayxiaodaiuserpermitAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_xiaodai_user_permit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
