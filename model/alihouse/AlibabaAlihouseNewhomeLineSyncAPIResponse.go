package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLineSyncAPIResponse 环线数据同步 API返回值
// alibaba.alihouse.newhome.line.sync
//
// 环线数据同步
type AlibabaAlihouseNewhomeLineSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeLineSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeLineSyncAPIResponseModel is 环线数据同步 成功返回结果
type AlibabaAlihouseNewhomeLineSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_line_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeLineSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
