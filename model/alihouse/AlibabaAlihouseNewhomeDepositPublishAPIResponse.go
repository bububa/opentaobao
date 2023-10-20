package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeDepositPublishAPIResponse 创建、修改、发布房产预存金商品 API返回值
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
type AlibabaAlihouseNewhomeDepositPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeDepositPublishAPIResponseModel
}

// AlibabaAlihouseNewhomeDepositPublishAPIResponseModel is 创建、修改、发布房产预存金商品 成功返回结果
type AlibabaAlihouseNewhomeDepositPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_deposit_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeDepositPublishResult `json:"result,omitempty" xml:"result,omitempty"`
}
