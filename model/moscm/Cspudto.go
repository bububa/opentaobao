package moscm

// Cspudto 
type Cspudto struct {

    // 货号
    
    ArtNo   string `json:"art_no,omitempty" xml:"art_no,omitempty"`
    

    // 条码
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // 创建时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 创建人唯一标识
    
    CreatorId   string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
    

    // 唯一标识
    
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    

    // 商品级别
    
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    

    // 修改时间
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

    // 修改人唯一标识
    
    ModifierId   string `json:"modifier_id,omitempty" xml:"modifier_id,omitempty"`
    

    // 外部商品Id
    
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    

    // 商品属性
    
    Properties   []PropertyDto `json:"properties,omitempty" xml:"properties,omitempty"`
    

    // spu
    
    Spu   *Spudto `json:"spu,omitempty" xml:"spu,omitempty"`
    

    // 是否标准产品
    
    Standard   bool `json:"standard,omitempty" xml:"standard,omitempty"`
    

    // 当前状态:删除(-1),正常(1）
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 款号
    
    StyleNo   string `json:"style_no,omitempty" xml:"style_no,omitempty"`
    

    // 子标题
    
    SubTitle   string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
    

    // 吊牌价,单位:元
    
    TagPrice   string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
    

    // 标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // userType
    
    UserType   string `json:"user_type,omitempty" xml:"user_type,omitempty"`
    

}
