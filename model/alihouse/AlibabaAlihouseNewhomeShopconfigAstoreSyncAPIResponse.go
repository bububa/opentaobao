package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse 天猫好房店铺装修-Astore上翻 API返回值
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
type AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel is 天猫好房店铺装修-Astore上翻 成功返回结果
type AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_astore_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *AlibabaAlihouseNewhomeShopconfigAstoreSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
