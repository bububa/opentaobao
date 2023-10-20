package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse 房源标准打标数据同步 API返回值
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
type AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel is 房源标准打标数据同步 成功返回结果
type AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_features_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseFeaturesSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
