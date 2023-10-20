package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuantaopasswordqueryAPIResponse 查询解析淘口令 API返回值
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
type AlibababaichuantaopasswordqueryAPIResponse struct {
	model.CommonResponse
	AlibababaichuantaopasswordqueryAPIResponseModel
}

// AlibababaichuantaopasswordqueryAPIResponseModel is 查询解析淘口令 成功返回结果
type AlibababaichuantaopasswordqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_taopassword_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BcTaoPasswordResult `json:"result,omitempty" xml:"result,omitempty"`
}
