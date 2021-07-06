package caipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoSignstatusCheckAPIResponse 检查用户是否签署支付宝代购协议 API返回值
// taobao.caipiao.signstatus.check
//
// 检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。
type TaobaoCaipiaoSignstatusCheckAPIResponse struct {
	model.CommonResponse
	TaobaoCaipiaoSignstatusCheckAPIResponseModel
}

// TaobaoCaipiaoSignstatusCheckAPIResponseModel is 检查用户是否签署支付宝代购协议 成功返回结果
type TaobaoCaipiaoSignstatusCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"caipiao_signstatus_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 签署支付宝代扣协议的Url
	SignUrl string `json:"sign_url,omitempty" xml:"sign_url,omitempty"`
	// 是否签署了支付宝代扣协议
	Sign bool `json:"sign,omitempty" xml:"sign,omitempty"`
}
