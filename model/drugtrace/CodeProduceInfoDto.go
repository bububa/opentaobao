package drugtrace

// CodeProduceInfoDto 
type CodeProduceInfoDto struct {

    // 码生产日期
    
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    

    // 码批次号
    
    ProduceBatchNo   string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
    

    // 码生产信息id
    
    ProduceInfoId   string `json:"produce_info_id,omitempty" xml:"produce_info_id,omitempty"`
    

    // 包装比例
    
    PkgRatio   string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
    

    // 生产信息集合
    
    ProduceInfoList   []ProduceInfoDto `json:"produce_info_list,omitempty" xml:"produce_info_list,omitempty"`
    

    // 产品序列码
    
    ProdSeqNo   string `json:"prod_seq_no,omitempty" xml:"prod_seq_no,omitempty"`
    

    // 企业唯一标识
    
    ProductRefEntId   string `json:"product_ref_ent_id,omitempty" xml:"product_ref_ent_id,omitempty"`
    

    // 企业主键ID
    
    ProductEntId   string `json:"product_ent_id,omitempty" xml:"product_ent_id,omitempty"`
    

    // 企业名称
    
    ProductEntName   string `json:"product_ent_name,omitempty" xml:"product_ent_name,omitempty"`
    

    // 制剂单位类型描述
    
    PrepnUnitDesc   string `json:"prepn_unit_desc,omitempty" xml:"prepn_unit_desc,omitempty"`
    

    // 制剂单位类型(详见码表)
    
    PrepnUnit   int64 `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
    

    // 最小制剂数量，单个
    
    PkgNum   int64 `json:"pkg_num,omitempty" xml:"pkg_num,omitempty"`
    

    // 最小制剂数量的总数
    
    SmallMeasureNum   int64 `json:"small_measure_num,omitempty" xml:"small_measure_num,omitempty"`
    

    // 该生产信息对应的最小包装数量
    
    SmallPkgNum   int64 `json:"small_pkg_num,omitempty" xml:"small_pkg_num,omitempty"`
    

    // 药品信息
    
    PhysicInfo   string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
    

    // 包装规格
    
    PkgSpec   string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
    

    // 药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类
    
    PhysicType   string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
    

    // 生产企业名称
    
    ProduceEntName   string `json:"produce_ent_name,omitempty" xml:"produce_ent_name,omitempty"`
    

    // 企业唯一标识
    
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    

    // 进口许可证
    
    LicenceNo   string `json:"licence_no,omitempty" xml:"licence_no,omitempty"`
    

    // 代理企业id
    
    AgentEntId   string `json:"agent_ent_id,omitempty" xml:"agent_ent_id,omitempty"`
    

    // 包装企业id
    
    PackEntId   string `json:"pack_ent_id,omitempty" xml:"pack_ent_id,omitempty"`
    

    // 生产企业id
    
    ProduceEntId   string `json:"produce_ent_id,omitempty" xml:"produce_ent_id,omitempty"`
    

    // 生产企业id
    
    ApprovalNo   string `json:"approval_no,omitempty" xml:"approval_no,omitempty"`
    

    // 授权企业id
    
    AuthrizerEntId   string `json:"authrizer_ent_id,omitempty" xml:"authrizer_ent_id,omitempty"`
    

    // 药品商品基本信息标识
    
    DrugEntBaseInfoId   string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
    

    // 当前企业id
    
    CurrEntId   string `json:"curr_ent_id,omitempty" xml:"curr_ent_id,omitempty"`
    

    // 有效期至
    
    ExprieDate   string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
    

}
