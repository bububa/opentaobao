package xhotelonlineorder

// HotelOrderDirectInfo 结构体
type HotelOrderDirectInfo struct {
	// 单次请求的唯一标识
	RequestID string `json:"request_i_d,omitempty" xml:"request_i_d,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 酒店编码
	HotelCode string `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// xml内容格式字符串，详细参考 http://open.taobao.com/docs/doc.htm?spm=a219a.7629140.0.0.2gmWOz&treeId=191&articleId=106152&docType=1
	Context string `json:"context,omitempty" xml:"context,omitempty"`
	// json格式的扩展字段
	Extensions string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// 给酒店前台展示的文案提示
	DisplayText string `json:"display_text,omitempty" xml:"display_text,omitempty"`
	// 淘宝订单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// * PMS处理消息的动作（pms端应该考虑增加代办任务列表）      * 0:普通，不需要特殊动作      * 1：弹窗提示（Pms中央模态窗）      * 2：托盘信息提示（给酒店广播通知或者平台公告）      * 3:发送请求给酒店前台人员，前台人员需要看到后确认处理(比如督促前台结账、督促前台上报入住状态)      * 4：请求pms自动处理(自动发起结账，自动上报订单状态等)
	Action int64 `json:"action,omitempty" xml:"action,omitempty"`
	// * 消息状态（tips:Pms定时get走请求，状态可以考虑不变化）      * 0：新建消息（接收到交易系统请求）      * 1:请求已获取（pms已经取走请求数据）      * 2:请求已认领（pms已经有人认领消息，正在处理）      * 3:请求已反馈（Pms反馈请求处理结果）      * 4:处理完成（已通知相关系统（交易））      * 5:请求失效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
