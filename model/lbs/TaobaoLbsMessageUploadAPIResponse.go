package lbs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLbsMessageUploadAPIResponse
lbs数据采集 API返回值
taobao.lbs.message.upload

lbs数据采集 */
type TaobaoLbsMessageUploadAPIResponse struct {
	model.CommonResponse
	TaobaoLbsMessageUploadAPIResponseModel
}

// TaobaoLbsMessageUploadAPIResponseModel is lbs数据采集 成功返回结果
type TaobaoLbsMessageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"lbs_message_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// subCode
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// subMsg
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
