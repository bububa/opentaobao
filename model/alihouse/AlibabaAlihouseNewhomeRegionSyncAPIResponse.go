package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeRegionSyncAPIResponse
城区数据同步 API返回值
alibaba.alihouse.newhome.region.sync

城区数据同步 */
type AlibabaAlihouseNewhomeRegionSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeRegionSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeRegionSyncAPIResponseModel is 城区数据同步 成功返回结果
type AlibabaAlihouseNewhomeRegionSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_region_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeRegionSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
