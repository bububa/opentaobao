package fundplatform

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// FundAccountResponse 结构体
type FundAccountResponse struct {
	// 外部订单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 账户名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 账户ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 可用余额
	BalanceAmount int64 `json:"balance_amount,omitempty" xml:"balance_amount,omitempty"`
	// 冻结金额
	FreezeAmount int64 `json:"freeze_amount,omitempty" xml:"freeze_amount,omitempty"`
	// 状态，1:正常,2:禁用
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
	// 用户ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolFundAccountResponse = sync.Pool{
	New: func() any {
		return new(FundAccountResponse)
	},
}

// GetFundAccountResponse() 从对象池中获取FundAccountResponse
func GetFundAccountResponse() *FundAccountResponse {
	return poolFundAccountResponse.Get().(*FundAccountResponse)
}

// ReleaseFundAccountResponse 释放FundAccountResponse
func ReleaseFundAccountResponse(v *FundAccountResponse) {
	v.OutBizId = ""
	v.Title = ""
	v.AccountId = 0
	v.BalanceAmount = 0
	v.FreezeAmount = 0
	v.Status = nil
	v.UserId = 0
	poolFundAccountResponse.Put(v)
}
