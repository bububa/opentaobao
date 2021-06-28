package wlb

// WlbItem 
/* model for simplify = false
type WlbItem struct {

    // 商品id
    
    Id   int64 `json:"id,omitempty"`
    

    // 商品所有人淘宝ID
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 商品所有人淘宝nick
    
    UserNick   string `json:"user_nick,omitempty"`
    

    // 商品的名称
    
    Name   string `json:"name,omitempty"`
    

    // 前台商品名称
    
    Title   string `json:"title,omitempty"`
    

    // 商家编码
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 是不是sku商品<br/>值为true或false
    
    IsSku   bool `json:"is_sku,omitempty"`
    

    // 标记，用逗号隔开的字符串。<br/>BIT_HAS_AUTHORIZE 第1位 是否有授权规则;<br/>BATCH  第2位 是否有批次规则；<br/>SYNCHRONIZATION 第3位 是否有同步规则。
    
    Flag   string `json:"flag,omitempty"`
    

    // 商品类型：<br/>NORMAL--普通类型;<br/>COMBINE--组合商品;<br/>DISTRIBUTION--分销商品;<br/>默认为NORMAL
    
    Type   string `json:"type,omitempty"`
    

    // 商品备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 父item的id，当item为物流宝子商品时，parent_id必填,否则不必填<br/>可通过父ID来得知商品的关系。
    
    ParentId   int64 `json:"parent_id,omitempty"`
    

    // 状态，item_status_valid -- 1 表示 有效 item_status_lock -- 2 表示锁住
    
    Status   string `json:"status,omitempty"`
    

    // 发布版本号，用来同步商
    
    PublishVersion   int64 `json:"publish_version,omitempty"`
    

    // 创建人
    
    Creator   string `json:"creator,omitempty"`
    

    // 品牌ID
    
    BrandId   int64 `json:"brand_id,omitempty"`
    

    // 创建日期
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 最后修改人
    
    LastModifier   string `json:"last_modifier,omitempty"`
    

    // 修改日期
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 属性key:value; key:value
    
    Properties   string `json:"properties,omitempty"`
    

    // 是否易碎
    
    IsFriable   bool `json:"is_friable,omitempty"`
    

    // 是否危险品
    
    IsDangerous   bool `json:"is_dangerous,omitempty"`
    

    // 颜色
    
    Color   string `json:"color,omitempty"`
    

    // 重量
    
    Weight   int64 `json:"weight,omitempty"`
    

    // mm
    
    Length   int64 `json:"length,omitempty"`
    

    // 宽
    
    Width   int64 `json:"width,omitempty"`
    

    // 高
    
    Height   int64 `json:"height,omitempty"`
    

    // 立方mm
    
    Volume   int64 `json:"volume,omitempty"`
    

    // 货类
    
    GoodsCat   string `json:"goods_cat,omitempty"`
    

    // 计价货类
    
    PricingCat   string `json:"pricing_cat,omitempty"`
    

    // 包装材料
    
    PackageMaterial   string `json:"package_material,omitempty"`
    

    // 价格
    
    Price   int64 `json:"price,omitempty"`
    

}
*/

// WlbItem 
type WlbItem struct {

    // 商品id
    Id   int64 `json:"id,omitempty"`

    // 商品所有人淘宝ID
    UserId   int64 `json:"user_id,omitempty"`

    // 商品所有人淘宝nick
    UserNick   string `json:"user_nick,omitempty"`

    // 商品的名称
    Name   string `json:"name,omitempty"`

    // 前台商品名称
    Title   string `json:"title,omitempty"`

    // 商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 是不是sku商品<br/>值为true或false
    IsSku   bool `json:"is_sku,omitempty"`

    // 标记，用逗号隔开的字符串。<br/>BIT_HAS_AUTHORIZE 第1位 是否有授权规则;<br/>BATCH  第2位 是否有批次规则；<br/>SYNCHRONIZATION 第3位 是否有同步规则。
    Flag   string `json:"flag,omitempty"`

    // 商品类型：<br/>NORMAL--普通类型;<br/>COMBINE--组合商品;<br/>DISTRIBUTION--分销商品;<br/>默认为NORMAL
    Type   string `json:"type,omitempty"`

    // 商品备注
    Remark   string `json:"remark,omitempty"`

    // 父item的id，当item为物流宝子商品时，parent_id必填,否则不必填<br/>可通过父ID来得知商品的关系。
    ParentId   int64 `json:"parent_id,omitempty"`

    // 状态，item_status_valid -- 1 表示 有效 item_status_lock -- 2 表示锁住
    Status   string `json:"status,omitempty"`

    // 发布版本号，用来同步商
    PublishVersion   int64 `json:"publish_version,omitempty"`

    // 创建人
    Creator   string `json:"creator,omitempty"`

    // 品牌ID
    BrandId   int64 `json:"brand_id,omitempty"`

    // 创建日期
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 最后修改人
    LastModifier   string `json:"last_modifier,omitempty"`

    // 修改日期
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 属性key:value; key:value
    Properties   string `json:"properties,omitempty"`

    // 是否易碎
    IsFriable   bool `json:"is_friable,omitempty"`

    // 是否危险品
    IsDangerous   bool `json:"is_dangerous,omitempty"`

    // 颜色
    Color   string `json:"color,omitempty"`

    // 重量
    Weight   int64 `json:"weight,omitempty"`

    // mm
    Length   int64 `json:"length,omitempty"`

    // 宽
    Width   int64 `json:"width,omitempty"`

    // 高
    Height   int64 `json:"height,omitempty"`

    // 立方mm
    Volume   int64 `json:"volume,omitempty"`

    // 货类
    GoodsCat   string `json:"goods_cat,omitempty"`

    // 计价货类
    PricingCat   string `json:"pricing_cat,omitempty"`

    // 包装材料
    PackageMaterial   string `json:"package_material,omitempty"`

    // 价格
    Price   int64 `json:"price,omitempty"`

}
