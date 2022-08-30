package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseStoreCheckAPIResponse 门店对账查询工具 API返回值
// alibaba.alihouse.store.check
//
// 门店对账查询工具
type AlibabaAlihouseStoreCheckAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseStoreCheckAPIResponseModel
}

// AlibabaAlihouseStoreCheckAPIResponseModel is 门店对账查询工具 成功返回结果
type AlibabaAlihouseStoreCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_store_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseStoreCheckResult `json:"result,omitempty" xml:"result,omitempty"`
}
