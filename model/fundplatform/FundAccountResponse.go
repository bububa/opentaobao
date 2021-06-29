package fundplatform
import (
    "github.com/bububa/opentaobao/model"
)

// FundAccountResponse 
type FundAccountResponse struct {
    // 账户ID
    AccountId   int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
    // 可用余额
    BalanceAmount   int64 `json:"balance_amount,omitempty" xml:"balance_amount,omitempty"`
    // 冻结金额
    FreezeAmount   int64 `json:"freeze_amount,omitempty" xml:"freeze_amount,omitempty"`
    // 外部订单号
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    // 状态，1:正常,2:禁用
    Status   *model.File `json:"status,omitempty" xml:"status,omitempty"`
    // 账户名称
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 用户ID
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
