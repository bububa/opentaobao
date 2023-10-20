package tblogistics

import (
	"sync"
)

// Shipping 结构体
type Shipping struct {
	// 拆单子订单列表，对应的数据是：该物流订单下的全部子订单
	SubTids []int64 `json:"sub_tids,omitempty" xml:"sub_tids>int64,omitempty"`
	// 物流订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 预约取货开始时间
	DeliveryStart string `json:"delivery_start,omitempty" xml:"delivery_start,omitempty"`
	// 预约取货结束时间
	DeliveryEnd string `json:"delivery_end,omitempty" xml:"delivery_end,omitempty"`
	// 运单号.具体一个物流公司的运单号码.
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 货物名称
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 收件人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 运单创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 运单修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 物流订单状态,可选值:CREATED(订单已创建) RECREATED(订单重新创建) CANCELLED(订单已取消) CLOSED(订单关闭) SENDING(等候发送给物流公司) ACCEPTING(已发送给物流公司,等待接单) ACCEPTED(物流公司已接单) REJECTED(物流公司不接单) PICK_UP(物流公司揽收成功) PICK_UP_FAILED(物流公司揽收失败) LOST(物流公司丢单) REJECTED_BY_RECEIVER(对方拒签) ACCEPTED_BY_RECEIVER(发货方式在线下单：对方已签收；自己联系：卖家已发货)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS).
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 谁承担运费.可选值:buyer(买家承担),seller(卖家承担运费).
	FreightPayer string `json:"freight_payer,omitempty" xml:"freight_payer,omitempty"`
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 卖家是否确认发货.可选值:yes(是),no(否).
	SellerConfirm string `json:"seller_confirm,omitempty" xml:"seller_confirm,omitempty"`
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 1111
	Openuid string `json:"openuid,omitempty" xml:"openuid,omitempty"`
	// 交易ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 表明是否是拆单，默认值0，1表示拆单
	IsSplit int64 `json:"is_split,omitempty" xml:"is_split,omitempty"`
	// 返回发货是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolShipping = sync.Pool{
	New: func() any {
		return new(Shipping)
	},
}

// GetShipping() 从对象池中获取Shipping
func GetShipping() *Shipping {
	return poolShipping.Get().(*Shipping)
}

// ReleaseShipping 释放Shipping
func ReleaseShipping(v *Shipping) {
	v.SubTids = v.SubTids[:0]
	v.OrderCode = ""
	v.SellerNick = ""
	v.BuyerNick = ""
	v.DeliveryStart = ""
	v.DeliveryEnd = ""
	v.OutSid = ""
	v.ItemTitle = ""
	v.ReceiverName = ""
	v.Created = ""
	v.Modified = ""
	v.Status = ""
	v.Type = ""
	v.FreightPayer = ""
	v.CompanyName = ""
	v.SellerConfirm = ""
	v.Ouid = ""
	v.Openuid = ""
	v.Tid = 0
	v.IsSplit = 0
	v.IsSuccess = false
	poolShipping.Put(v)
}
