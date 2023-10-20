package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxAccountAccountQueryAPIResponse 万相台账号余额查询 API返回值
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
type TaobaoOnebpDkxAccountAccountQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxAccountAccountQueryAPIResponseModel
}

// TaobaoOnebpDkxAccountAccountQueryAPIResponseModel is 万相台账号余额查询 成功返回结果
type TaobaoOnebpDkxAccountAccountQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_account_account_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
