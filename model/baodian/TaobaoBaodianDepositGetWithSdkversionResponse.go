package baodian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户宝点信息（带sdk版本，已迁移） APIResponse
taobao.baodian.deposit.get.with.sdkversion

获取用户宝点信息（带sdk版本，已迁移）
*/
type TaobaoBaodianDepositGetWithSdkversionAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"baodian_deposit_get_with_sdkversion_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结构体
    
    Result   *CoinUserDepositV2 `json:"result,omitempty" xml:"