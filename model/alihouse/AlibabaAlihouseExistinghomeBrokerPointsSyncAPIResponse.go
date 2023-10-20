package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse 经纪人积分同步 API返回值
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
type AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel
}

// AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel is 经纪人积分同步 成功返回结果
type AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_broker_points_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBrokerPointsSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
