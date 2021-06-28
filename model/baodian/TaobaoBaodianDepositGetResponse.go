package baodian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
宝点用户帐户查询（已迁移） APIResponse
taobao.baodian.deposit.get

查询用户宝点帐户信息及当前宝点价格
*/
type TaobaoBaodianDepositGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaodianDepositGetResponse `json:"baodian_deposit_get_response,omitempty"` 
    TaobaoBaodianDepositGetResponse
}

/* model for simplify = false
type TaobaoBaodianDepositGetResponse struct {

    // 用户宝点帐户信息
    
    UserCoinDeposit  *struct {
        UserCoinDeposit  *UserCoinDeposit `json:"user_coin_deposit,omitempty"`
    } `json:"user_coin_deposit,omitempty"`
    

}
*/

type TaobaoBaodianDepositGetResponse struct {

    // 用户宝点帐户信息
    UserCoinDeposit   *UserCoinDeposit `json:"user_coin_deposit,omitempty"`

}
