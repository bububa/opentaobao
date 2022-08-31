package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseSyncAPIResponse 房源信息同步 API返回值
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
type AlibabaAlihouseExistinghomeHouseSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel
}

// AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel is 房源信息同步 成功返回结果
type AlibabaAlihouseExistinghomeHouseSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_house_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeHouseSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
