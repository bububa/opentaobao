package baodian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
宝点用户帐户查询（已迁移） API返回值 
taobao.baodian.deposit.get

查询用户宝点帐户信息及当前宝点价格
*/
type TaobaoBaodianDepositGetAPIResponse struct {
    model.CommonResponse
    TaobaoBaodianDepositGetAPIResponseModel
}

// 宝点用户帐户查询（已迁移） 成功返回结果
type TaobaoBaodianDepositGetAPIResponseModel struct {
    XMLName xml.Name `xml:"baodian_deposit_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 用户宝点帐户信息
    UserCoinDeposit   *UserCoinDeposit `json:"user_coin_deposit,omitempty" xml:"user_coin_deposit,omitempty"`
}
