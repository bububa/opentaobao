package fenxiao

// ScItem 
type ScItem struct {

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    

    // 商家编码
    
    OuterCode   string `json:"outer_code,omitempty" xml:"outer_code,omitempty"`
    

    // 1.普通供应链商品 2.供应链组合主商品
    
    ItemType   int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
    

    // 商品属性格式是  p1:v1,p2:v2,p3:v3
    
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
    

    // 条形码
    
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    

    // 仓储商编码，可以支持多个，格式wmsname:code
    
    WmsCode   string `json:"wms_code,omitempty" xml:"wms_code,omitempty"`
    

    // 是否易碎 false ：不是  true：是
    
    IsFriable   bool `json:"is_friable,omitempty" xml:"is_friable,omitempty"`
    

    // 是否危险 false：不是  true：是
    
    IsDangerous   bool `json:"is_dangerous,omitempty" xml:"is_dangerous,omitempty"`
    

    // 贵重品:false:不是 true：是
    
    IsCostly   bool `json:"is_costly,omitempty" xml:"is_costly,omitempty"`
    

    // 是否保质期：false:不是 true：是
    
    IsWarranty   bool `json:"is_warranty,omitempty" xml:"is_warranty,omitempty"`
    

    // 重量.单位：克
    
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    

    // 长度 单位：mm
    
    Length   int64 `json:"length,omitempty" xml:"length,omitempty"`
    

    // 宽 单位：mm
    
    Width   int64 `json:"width,omitempty" xml:"width,omitempty"`
    

    // 高 单位：mm
    
    Height   int64 `json:"height,omitempty" xml:"height,omitempty"`
    

    // 体积：立方厘米
    
    Volume   int64 `json:"volume,omitempty" xml:"volume,omitempty"`
    

    // 价格：分（吊牌价）
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // LIQUID:液体，1：粉体，SOLID：固体
    
    MatterStatus   string `json:"matter_status,omitempty" xml:"matter_status,omitempty"`
    

    // 品牌id
    
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 品牌名称
    
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    

    // 1表示区域销售，0或是空是非区域销售
    
    IsAreaSale   int64 `json:"is_area_sale,omitempty" xml:"is_area_sale,omitempty"`
    

    // 后端商品options字段
    
    Options   int64 `json:"options,omitempty" xml:"options,omitempty"`
    

}
