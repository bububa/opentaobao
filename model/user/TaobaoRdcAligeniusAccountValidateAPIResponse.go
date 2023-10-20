package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniusaccountvalidateAPIResponse AG商家账号校验 API返回值
// taobao.rdc.aligenius.account.validate
//
// 提供应对接AG的erp系统查询其旗下的商家是否为AG商家
type TaobaordcaligeniusaccountvalidateAPIResponse struct {
	model.CommonResponse
	TaobaordcaligeniusaccountvalidateAPIResponseModel
}

// TaobaordcaligeniusaccountvalidateAPIResponseModel is AG商家账号校验 成功返回结果
type TaobaordcaligeniusaccountvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_account_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaordcaligeniusaccountvalidateResult `json:"result,omitempty" xml:"result,omitempty"`
}
