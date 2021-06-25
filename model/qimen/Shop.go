package qimen

// Shop 
type Shop struct {

    // 来源平台编码, string (50) , 必填,TB= 淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、MIA=蜜芽、GW=商家官网、CT=村淘、YJWD=云集微店、OTHERS=其他,  (只传英文编码)
    SourcePlatformCode   string `json:"sourcePlatformCode,omitempty"`

    // 货主编码, string (50) , 必填
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 仓库编码, string (50)，必填
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // ERP店铺编码,  string (50) ,  必填
    ShopCode   string `json:"shopCode,omitempty"`

    // ERP店铺名称，string (50)
    ShopName   string `json:"shopName,omitempty"`

    // 平台店铺Id（如淘宝店铺Id）,  string (50)
    PlatformShopCode   string `json:"platformShopCode,omitempty"`

    // 平台店铺名称（如淘宝店铺名称）,  string (50)
    PlatformShopName   string `json:"platformShopName,omitempty"`

    // 店铺地址
    ShopAddress   *ShopAddress `json:"shopAddress,omitempty"`

}
