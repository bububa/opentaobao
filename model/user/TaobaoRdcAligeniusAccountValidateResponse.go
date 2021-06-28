package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AG商家账号校验 APIResponse
taobao.rdc.aligenius.account.validate

提供应对接AG的erp系统查询其旗下的商家是否为AG商家
*/
type TaobaoRdcAligeniusAccountValidateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rdc_aligenius_account_validate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdcAligeniusAccountValidateResult `json:"result,omitempty" xml:"