package wdk

// ItemInfo 
type ItemInfo struct {
    // 商品状态：商品在机构内的生命周期，商品状态；A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡，传ATCRD
    LifeStatus   string `json:"life_status,omitempty" xml:"life_status,omitempty"`
    // 商品名称：商品名称；显示在手机APP商品详情页，对商品直观的描述，通常包含了品牌、规格等信息。商品名称也会显示在pos小票上。附：淘鲜达系统的商品品名禁止出现“专供”，“特供”，“聚划算”字样。
    SkuName   string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
    // 商品类型：商品类型类型 1:普通商品、2:加工半成品、3:加工成品、4:原材料、5:耗材
    ItemTypeNew   string `json:"item_type_new,omitempty" xml:"item_type_new,omitempty"`
    // 商品编码：商品编码（此字段不能修改），商品在商家的唯一编码，传商家本地ERP一致的商品编码（货号）
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 商家类目编码：商家类目code，可以为空
    MerchantCatCode   string `json:"merchant_cat_code,omitempty" xml:"merchant_cat_code,omitempty"`
    // 库存单位：库存单位（此字段一经录入不能修改），存储的最小单位，从计量单位表中选择，如果没有对应的单位，可以由运营联系淘鲜达的运营新增单位。填写下列单位选项：毫升、升、加仑、屉、批、面、刀、幅、册、口、床、捆、壶、版、平方米、米、厘米、毫米、笼、坛、节、筒、半只、卡、板、双、粒、令、筐、碗、排、盘、顶、本、封、颗、付、辆、台、次、张、g、盆、组、扎、只、支、提、套、束、片、枚、篮、棵、卷、听、罐、根、副、份、朵、对、袋、打、串、杯、把、kg、件、箱、块、条、瓶、桶、包、盒、个。若新增的单位需要支持小数点库存，联系飞观。手机淘宝中淘鲜达价格展示是用的库存单位。商家在填写商品库存单位时，应注意此单位展示合理性；在库存和销售单位实际不一致的情况（例如某类商品库存使用kg，销售使用g）之外，建议商品库存单位和销售单位一致。
    InventoryUtil   string `json:"inventory_util,omitempty" xml:"inventory_util,omitempty"`
    // 机构名称：机构名称，由淘鲜达商家运营分配
    OrgName   string `json:"org_name,omitempty" xml:"org_name,omitempty"`
    // 机构编码：机构编码，由淘鲜达商家运营分配
    OrgCode   string `json:"org_code,omitempty" xml:"org_code,omitempty"`
    // 盒马类目ID：需要商家把自己的类目对应到盒马的类目上，填写盒马类目ID，如果商家不填写，系统则会进行类目预测，将预测到的类目填写进去
    BackCatCode   string `json:"back_cat_code,omitempty" xml:"back_cat_code,omitempty"`
    // 产地：需要传淘鲜达产地库中的值；国内产地传值格式：中国|省|市。若不能确定产地，可以传“见产品外包装”。国外产地只需要传国家名
    ProducerPlace   string `json:"producer_place,omitempty" xml:"producer_place,omitempty"`
    // 供应商ID：供应商ID；淘鲜达合作商家填虚拟供应商ID，一共2个（自营、联营），由淘鲜达运营帮忙创建，联营商品需填写联营供应商ID，该商品可以通过联营供应商RF枪扫码入库库存
    SupplierNo   string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
    // 销售规格描述：商品销售单位对应的含量表达；APP展示重要字段，体现售卖单位中含有的商品数量，通常描述为“550g/份”等样式。填字符串。
    SaleSpec   string `json:"sale_spec,omitempty" xml:"sale_spec,omitempty"`
    // 商品简称：对于商品的描述进行简化的表达，以便在拣货、标签等页面上能够进行直接展示；最多40个字符，会显示在pos屏幕上
    ShortTitle   string `json:"short_title,omitempty" xml:"short_title,omitempty"`
    // 门店商品售价：商品一个库存单位的售卖价格，2位小数
    SkuPrice   int64 `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
    // 门店商品会员售价：会员正常购买该商品的售价，2位小数
    MemberPrice   int64 `json:"member_price,omitempty" xml:"member_price,omitempty"`
    // 重量
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    // 是否进口：是否进口,是否进口；原产地非中国，都填是 1：是? 0：否（默认为0）
    ImportFlag   int64 `json:"import_flag,omitempty" xml:"import_flag,omitempty"`
    // 品牌名称：商品的品牌名字，需要先在淘鲜达提供的品牌库中匹配，如果匹配不上的直接传ERP中的品牌值
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 净含量：商品包装规格的描述，建议跟销售规格描述填写一致
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 存储条件：存储条件；填常温、冷藏、冷冻、热链、鲜活
    Storage   string `json:"storage,omitempty" xml:"storage,omitempty"`
    // 保质天数：商品的保质期天数，必须为整数，0代表不管理保质期
    Period   int64 `json:"period,omitempty" xml:"period,omitempty"`
    // 是否称重：是否称重（此字段一经录入不能修改）；若库存单位是kg，或g，此字段填是，其他的填否 1：是? 0：否（默认为0)
    WeightFlag   int64 `json:"weight_flag,omitempty" xml:"weight_flag,omitempty"`
    // 条码：包含商品包装上已经印刷的条码，也包含企业内部的自编条码用于打印在食品包装上
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    // 标价签类型：标价签类型；商品在门店陈列时，采用的售价标签类型。因不采用电子价签，建议商家填默认值“无价签”
    LabelStyleType   string `json:"label_style_type,omitempty" xml:"label_style_type,omitempty"`
    // 售卖单位
    SaleUtil   string `json:"sale_util,omitempty" xml:"sale_util,omitempty"`
    // 均重
    AvgWeight   int64 `json:"avg_weight,omitempty" xml:"avg_weight,omitempty"`
}
