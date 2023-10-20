package idle

import (
	"sync"
)

// RentalOrderDto 结构体
type RentalOrderDto struct {
	// 订单商品信息
	Items []RentalOrderItemDto `json:"items,omitempty" xml:"items>rental_order_item_dto,omitempty"`
	// 买家id
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 预约时间
	ReservedPackageTime string `json:"reserved_package_time,omitempty" xml:"reserved_package_time,omitempty"`
	// 用户会员卡到期时间
	BuyerMemberExpireDate string `json:"buyer_member_expire_date,omitempty" xml:"buyer_member_expire_date,omitempty"`
	// 买家地址
	BuyerAddress *UserAddressDto `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 邮费，单位分
	Postage int64 `json:"postage,omitempty" xml:"postage,omitempty"`
}

var poolRentalOrderDto = sync.Pool{
	New: func() any {
		return new(RentalOrderDto)
	},
}

// GetRentalOrderDto() 从对象池中获取RentalOrderDto
func GetRentalOrderDto() *RentalOrderDto {
	return poolRentalOrderDto.Get().(*RentalOrderDto)
}

// ReleaseRentalOrderDto 释放RentalOrderDto
func ReleaseRentalOrderDto(v *RentalOrderDto) {
	v.Items = v.Items[:0]
	v.BuyerId = ""
	v.ReservedPackageTime = ""
	v.BuyerMemberExpireDate = ""
	v.BuyerAddress = nil
	v.OrderId = 0
	v.Status = 0
	v.Postage = 0
	poolRentalOrderDto.Put(v)
}
