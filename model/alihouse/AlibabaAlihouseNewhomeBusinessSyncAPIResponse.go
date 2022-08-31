package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeBusinessSyncAPIResponse 商圈数据同步 API返回值
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
type AlibabaAlihouseNewhomeBusinessSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel is 商圈数据同步 成功返回结果
type AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_business_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeBusinessSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
