package axintrade

import (
	"sync"
)

// AxinFundConfirmResDto 结构体
type AxinFundConfirmResDto struct {
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 结算金额
	SettleAmount int64 `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
}

var poolAxinFundConfirmResDto = sync.Pool{
	New: func() any {
		return new(AxinFundConfirmResDto)
	},
}

// GetAxinFundConfirmResDto() 从对象池中获取AxinFundConfirmResDto
func GetAxinFundConfirmResDto() *AxinFundConfirmResDto {
	return poolAxinFundConfirmResDto.Get().(*AxinFundConfirmResDto)
}

// ReleaseAxinFundConfirmResDto 释放AxinFundConfirmResDto
func ReleaseAxinFundConfirmResDto(v *AxinFundConfirmResDto) {
	v.OuterOrderId = ""
	v.SettleAmount = 0
	poolAxinFundConfirmResDto.Put(v)
}
