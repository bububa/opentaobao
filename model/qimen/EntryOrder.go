package qimen

// EntryOrder 结构体
type EntryOrder struct {
	// 订单编码
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 后端订单id
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 订单类型
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 仓库名称
	WarehouseName string `json:"warehouseName,omitempty" xml:"warehouseName,omitempty"`
	// 单据总行数(当单据需要分多个请求发送时;发送方需要将totalOrderLines填入;接收方收到后;根据实际接收到的 条数和 totalOrderLines进行比对;如果小于;则继续等待接收请求。如果等于;则表示该单据的所有请求发送完 成)
	TotalOrderLines int64 `json:"totalOrderLines,omitempty" xml:"totalOrderLines,omitempty"`
	// 入库单号
	EntryOrderCode string `json:"entryOrderCode,omitempty" xml:"entryOrderCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 采购单号(当orderType=CGRK时使用)
	PurchaseOrderCode string `json:"purchaseOrderCode,omitempty" xml:"purchaseOrderCode,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 仓储系统入库单ID
	EntryOrderId string `json:"entryOrderId,omitempty" xml:"entryOrderId,omitempty"`
	// 入库单类型(SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入 库;B2BRK=B2B入库)
	EntryOrderType string `json:"entryOrderType,omitempty" xml:"entryOrderType,omitempty"`
	// 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求 不会被重复处理)
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 支持出入库单多次收货(多次收货后确认时:0:表示入库单最终状态确认;1:表示入库单中间状态确认;每次入库传入的数量为增 量;特殊情况;同一入库单;如果先收到0;后又收到1;允许修改收货的数量)
	ConfirmType int64 `json:"confirmType,omitempty" xml:"confirmType,omitempty"`
	// 入库单状态(NEW-未开始处理;ACCEPT-仓库接单;PARTFULFILLED-部分收货完成;FULFILLED-收货完成;EXCEPTION-异 常;CANCELED-取消;CLOSED-关闭;REJECT-拒单;CANCELEDFAIL-取消失败;只传英文编码)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 操作时间(YYYY-MM-DD HH:MM:SS;当status=FULFILLED;operateTime为入库时间)
	OperateTime string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 邮费
	Freight string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 入库单确认的其他入库子类型，用于entryOrderType设置为其他入库时设置
	SubOrderType string `json:"subOrderType,omitempty" xml:"subOrderType,omitempty"`
	// 该笔入库单的费用承担部门或责任部门
	ResponsibleDepartment string `json:"responsibleDepartment,omitempty" xml:"responsibleDepartment,omitempty"`
	// 店铺名称
	ShopNick string `json:"shopNick,omitempty" xml:"shopNick,omitempty"`
	// 店铺编码
	ShopCode string `json:"shopCode,omitempty" xml:"shopCode,omitempty"`
	// 订单创建时间(YYYY-MM-DD HH:MM:SS)
	OrderCreateTime string `json:"orderCreateTime,omitempty" xml:"orderCreateTime,omitempty"`
	// 关联订单信息
	RelatedOrders []RelatedOrder `json:"relatedOrders,omitempty" xml:"relatedOrders>related_order,omitempty"`
	// 预期到货时间(YYYY-MM-DD HH:MM:SS)
	ExpectStartTime string `json:"expectStartTime,omitempty" xml:"expectStartTime,omitempty"`
	// 最迟预期到货时间(YYYY-MM-DD HH:MM:SS)
	ExpectEndTime string `json:"expectEndTime,omitempty" xml:"expectEndTime,omitempty"`
	// 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中通(ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他(只传英文编码))
	LogisticsCode string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
	// 物流公司名称
	LogisticsName string `json:"logisticsName,omitempty" xml:"logisticsName,omitempty"`
	// 运单号
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
	// 操作员编码
	OperatorCode string `json:"operatorCode,omitempty" xml:"operatorCode,omitempty"`
	// 操作员名称
	OperatorName string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// 发件人信息
	SenderInfo *SenderInfo `json:"senderInfo,omitempty" xml:"senderInfo,omitempty"`
	// 收件人信息
	ReceiverInfo *ReceiverInfo `json:"receiverInfo,omitempty" xml:"receiverInfo,omitempty"`
	// 出库仓库编码
	SourceWarehouseCode string `json:"sourceWarehouseCode,omitempty" xml:"sourceWarehouseCode,omitempty"`
	// 出库仓库名称
	SourceWarehouseName string `json:"sourceWarehouseName,omitempty" xml:"sourceWarehouseName,omitempty"`
}
