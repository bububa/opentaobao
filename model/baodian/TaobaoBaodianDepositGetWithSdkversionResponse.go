package baodian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询用户宝点信息（带sdk版本，已迁移） APIResponse
taobao.baodian.deposit.get.with.sdkversion

获取用户宝点信息（带sdk版本，已迁移）
*/
type TaobaoBaodianDepositGetWithSdkversionAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaodianDepositGetWithSdkversionResponse `json:"baodian_deposit_get_with_sdkversion_response,omitempty"` 
    TaobaoBaodianDepositGetWithSdkversionResponse
}

/* model for simplify = false
type TaobaoBaodianDepositGetWithSdkversionResponse struct {

    // 结构体
    
    Result  *struct {
        CoinUserDepositV2  *CoinUserDepositV2 `json:"coin_user_deposit_v2,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBaodianDepositGetWithSdkversionResponse struct {

    // 结构体
    Result   *CoinUserDepositV2 `json:"result,omitempty"`

}
