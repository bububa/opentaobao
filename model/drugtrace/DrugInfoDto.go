package drugtrace

// DrugInfoDto 
type DrugInfoDto struct {

    // 20位码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 药品通用名
    
    DrugName   string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
    

    // 制剂规格
    
    Specifications   string `json:"specifications,omitempty" xml:"specifications,omitempty"`
    

    // 11：没有申请过此码 12：后四位检验不对 13：没激活过
    
    SubRetCode   string `json:"sub_ret_code,omitempty" xml:"sub_ret_code,omitempty"`
    

    // 状态码   0:正常  -1：未知异常  1：有原因的异常 见subretcode
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    

    // 剂型
    
    PrepnType   string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
    

    // 包装规格
    
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    

    // 查询次数(用户维度)
    
    QueryTotalCount   int64 `json:"query_total_count,omitempty" xml:"query_total_count,omitempty"`
    

    // 批次
    
    ProductionBatch   string `json:"production_batch,omitempty" xml:"production_batch,omitempty"`
    

    // 扫码次数
    
    QueryCount   string `json:"query_count,omitempty" xml:"query_count,omitempty"`
    

    // 是否销售
    
    IsSale   string `json:"is_sale,omitempty" xml:"is_sale,omitempty"`
    

    // 销售时间
    
    SaleTime   string `json:"sale_time,omitempty" xml:"sale_time,omitempty"`
    

    // 生产企业名称
    
    EntName   string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
    

    // 有效期至
    
    ExpiryDate   string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
    

    // 第一次查询时间
    
    FirstQuery   string `json:"first_query,omitempty" xml:"first_query,omitempty"`
    

    // 生产日期
    
    ProductionDate   string `json:"production_date,omitempty" xml:"production_date,omitempty"`
    

    // 销售企业
    
    SaleEnt   string `json:"sale_ent,omitempty" xml:"sale_ent,omitempty"`
    

    // 追溯码状态
    
    CodeStatus   string `json:"code_status,omitempty" xml:"code_status,omitempty"`
    

}
