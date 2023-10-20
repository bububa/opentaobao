package wlb

import (
	"sync"
)

// WlbOrderDetail 结构体
type WlbOrderDetail struct {
	// 物流宝订单商品
	OrderItemList []WlbOrderItem `json:"order_item_list,omitempty" xml:"order_item_list>wlb_order_item,omitempty"`
	// 出库或者入库，IN表示入库，OUT表示出库
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 订单来源:&lt;br/&gt;产生物流订单的原因，比如:&lt;br/&gt;&lt;br/&gt;订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用
	OrderSourceCode string `json:"order_source_code,omitempty" xml:"order_source_code,omitempty"`
	// 1:正常订单: NARMAL&lt;br/&gt;2:退货订单: RETURN&lt;br/&gt;3:换货订单: CHANGE
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// (1)其它:    OTHER&lt;br/&gt;(2)淘宝交易: TAOBAO&lt;br/&gt;(3)301:调拨: ALLOCATION&lt;br/&gt;(4)401:盘点:CHECK&lt;br/&gt;(5)501:销售采购:PRUCHASE
	OrderSubType string `json:"order_sub_type,omitempty" xml:"order_sub_type,omitempty"`
	// 卖家NICK
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 物流状态，&lt;br/&gt;订单已创建：0&lt;br/&gt;订单已取消: -1&lt;br/&gt;订单关闭:-3&lt;br/&gt;下发中: 10&lt;br/&gt;已下发: 11&lt;br/&gt;接收方拒签 :-100&lt;br/&gt;已发货:100&lt;br/&gt;签收成功:200
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 如果是交易单，则显示交易中买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单创建时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 订单最后一次修改时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 买家openId
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 卖家ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolWlbOrderDetail = sync.Pool{
	New: func() any {
		return new(WlbOrderDetail)
	},
}

// GetWlbOrderDetail() 从对象池中获取WlbOrderDetail
func GetWlbOrderDetail() *WlbOrderDetail {
	return poolWlbOrderDetail.Get().(*WlbOrderDetail)
}

// ReleaseWlbOrderDetail 释放WlbOrderDetail
func ReleaseWlbOrderDetail(v *WlbOrderDetail) {
	v.OrderItemList = v.OrderItemList[:0]
	v.OperateType = ""
	v.OrderCode = ""
	v.OrderSource = ""
	v.OrderSourceCode = ""
	v.OrderType = ""
	v.OrderSubType = ""
	v.UserNick = ""
	v.StoreCode = ""
	v.OrderStatus = ""
	v.Remark = ""
	v.BuyerNick = ""
	v.ModifyTime = ""
	v.CreateTime = ""
	v.OpenUid = ""
	v.UserId = 0
	poolWlbOrderDetail.Put(v)
}
