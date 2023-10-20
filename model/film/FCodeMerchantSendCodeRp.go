package film

import (
	"sync"
)

// FCodeMerchantSendCodeRp 结构体
type FCodeMerchantSendCodeRp struct {
	// 码对外信息描述列表
	FCodeMerchantInfoList []FCodeMerchantVo `json:"f_code_merchant_info_list,omitempty" xml:"f_code_merchant_info_list>f_code_merchant_vo,omitempty"`
}

var poolFCodeMerchantSendCodeRp = sync.Pool{
	New: func() any {
		return new(FCodeMerchantSendCodeRp)
	},
}

// GetFCodeMerchantSendCodeRp() 从对象池中获取FCodeMerchantSendCodeRp
func GetFCodeMerchantSendCodeRp() *FCodeMerchantSendCodeRp {
	return poolFCodeMerchantSendCodeRp.Get().(*FCodeMerchantSendCodeRp)
}

// ReleaseFCodeMerchantSendCodeRp 释放FCodeMerchantSendCodeRp
func ReleaseFCodeMerchantSendCodeRp(v *FCodeMerchantSendCodeRp) {
	v.FCodeMerchantInfoList = v.FCodeMerchantInfoList[:0]
	poolFCodeMerchantSendCodeRp.Put(v)
}
