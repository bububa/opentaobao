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
    Response *TaobaoBaodianDepositGetWithSdkversionResponse `json:"taobao_baodian_deposit_get_with_sdkversion_response,omitempty"`
}

type TaobaoBaodianDepositGetWithSdkversionResponse struct {

    // 结构体
    Result   *CoinUserDepositV2 `json:"result,omitempty"`

}
