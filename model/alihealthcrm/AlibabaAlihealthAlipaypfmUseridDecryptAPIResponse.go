package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmUseridDecryptAPIResponse
用户标识解密接口 API返回值
alibaba.alihealth.alipaypfm.userid.decrypt

用户唯一表示加密传输，调用方解密 */
type AlibabaAlihealthAlipaypfmUseridDecryptAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlipaypfmUseridDecryptAPIResponseModel
}

// AlibabaAlihealthAlipaypfmUseridDecryptAPIResponseModel is 用户标识解密接口 成功返回结果
type AlibabaAlihealthAlipaypfmUseridDecryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_userid_decrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
