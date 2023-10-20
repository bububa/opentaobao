package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomedepositpublishAPIResponse 创建、修改、发布房产预存金商品 API返回值
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
type AlibabaalihousenewhomedepositpublishAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomedepositpublishAPIResponseModel
}

// AlibabaalihousenewhomedepositpublishAPIResponseModel is 创建、修改、发布房产预存金商品 成功返回结果
type AlibabaalihousenewhomedepositpublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_deposit_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomedepositpublishResult `json:"result,omitempty" xml:"result,omitempty"`
}
