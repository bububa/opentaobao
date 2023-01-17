package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseStorePunishInfoQueryAPIResponse 门店处罚信息查询 API返回值
// alibaba.alihouse.store.punish.info.query
//
// 门店处罚信息查询
type AlibabaAlihouseStorePunishInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseStorePunishInfoQueryAPIResponseModel
}

// AlibabaAlihouseStorePunishInfoQueryAPIResponseModel is 门店处罚信息查询 成功返回结果
type AlibabaAlihouseStorePunishInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_store_punish_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseStorePunishInfoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
