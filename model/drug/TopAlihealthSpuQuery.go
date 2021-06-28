package drug

// TopAlihealthSpuQuery 
/* model for simplify = false
type TopAlihealthSpuQuery struct {

    // 69码
    
    Barcode   string `json:"barcode,omitempty"`
    

    // 批准文号
    
    ApprovalNumber   string `json:"approval_number,omitempty"`
    

    // 通用名称
    
    Title   string `json:"title,omitempty"`
    

    // 规格
    
    Specification   string `json:"specification,omitempty"`
    

    // 生产厂商
    
    Manufacturer   string `json:"manufacturer,omitempty"`
    

    // 品牌名称
    
    BrandName   string `json:"brand_name,omitempty"`
    

    // 是否只搜索药品
    
    OnlyDrug   bool `json:"only_drug,omitempty"`
    

    // 开放id
    
    OpenId   int64 `json:"open_id,omitempty"`
    

    // 偏移量
    
    Offset   int64 `json:"offset,omitempty"`
    

    // 当前页码
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 每页条数，0不分页，返回全部
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 药品SPUID
    
    SpuId   int64 `json:"spu_id,omitempty"`
    

    // spu标题
    
    ProductTitle   string `json:"product_title,omitempty"`
    

    // 别名，也就是商品名称
    
    TitleAlias   string `json:"title_alias,omitempty"`
    

    // 药物说明书信息
    
    AlihealthDrugInstructionDTO  *struct {
        TopAlihealthDrugInstructionDto  *TopAlihealthDrugInstructionDto `json:"top_alihealth_drug_instruction_dto,omitempty"`
    } `json:"alihealth_drug_instruction_d_t_o,omitempty"`
    

}
*/

// TopAlihealthSpuQuery 
type TopAlihealthSpuQuery struct {

    // 69码
    Barcode   string `json:"barcode,omitempty"`

    // 批准文号
    ApprovalNumber   string `json:"approval_number,omitempty"`

    // 通用名称
    Title   string `json:"title,omitempty"`

    // 规格
    Specification   string `json:"specification,omitempty"`

    // 生产厂商
    Manufacturer   string `json:"manufacturer,omitempty"`

    // 品牌名称
    BrandName   string `json:"brand_name,omitempty"`

    // 是否只搜索药品
    OnlyDrug   bool `json:"only_drug,omitempty"`

    // 开放id
    OpenId   int64 `json:"open_id,omitempty"`

    // 偏移量
    Offset   int64 `json:"offset,omitempty"`

    // 当前页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 每页条数，0不分页，返回全部
    PageSize   int64 `json:"page_size,omitempty"`

    // 药品SPUID
    SpuId   int64 `json:"spu_id,omitempty"`

    // spu标题
    ProductTitle   string `json:"product_title,omitempty"`

    // 别名，也就是商品名称
    TitleAlias   string `json:"title_alias,omitempty"`

    // 药物说明书信息
    AlihealthDrugInstructionDTO   *TopAlihealthDrugInstructionDto `json:"alihealth_drug_instruction_d_t_o,omitempty"`

}
