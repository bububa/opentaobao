package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBankcardSyncAPIResponse 结算账号同步 API返回值
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
type AlibabaAlihouseExistinghomeBankcardSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel
}

// AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel is 结算账号同步 成功返回结果
type AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_bankcard_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBankcardSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
