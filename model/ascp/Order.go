package ascp

// Order 结构体
type Order struct {
	// ERP单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// WMS单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 单据类型, string (50) ,必填 , JYCK= 一般交易出库单，HHCK= 换货出库 ，BFCK= 补发出库，PTCK=普通出库单，DBCK=调拨出库 ，B2BRK=B2B入库，B2BCK=B2B出库，QTCK=其他出库，SCRK=生产入库，LYRK=领用入库，CCRK=残次品入库，CGRK=采购入库 ，DBRK= 调拨入库 ，QTRK= 其他入库 ，XTRK= 销退入库，HHRK= 换货入库，CNJG= 仓内加工单，BIGTOBCK=大B2B出库
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 订单来源平台编码(TB= 淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、TGC=淘工厂、POS=POS门店、MIA=蜜芽、GW=商家官网、CT=村淘、YJWD=云集微店、PDD=拼多多、YZ=有赞、DY=抖音、KS=快手、TXP=淘小铺、FX=分销、XHS=小红书、DW=得物、DWZF=得物直发、TMGJ=天猫国际、TMCS=天猫超市、OTHERS=其他)
	SourcePlatformCode string `json:"source_platform_code,omitempty" xml:"source_platform_code,omitempty"`
	// 订单来源平台名称
	SourcePlatformName string `json:"source_platform_name,omitempty" xml:"source_platform_name,omitempty"`
	// 店铺名称
	ShopNick string `json:"shop_nick,omitempty" xml:"shop_nick,omitempty"`
	// ERP仓库编码,sellerId下唯一
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
	// 创建时间（时间戳）
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}
