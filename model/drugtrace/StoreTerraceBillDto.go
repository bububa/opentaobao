package drugtrace

// StoreTerraceBillDto 
type StoreTerraceBillDto struct {

    // 单据ID
    
    BillId   string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
    

    // 单据编码
    
    BillCode   string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
    

    // 企业ID
    
    EntId   string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
    

    // 创建时间
    
    CrtDate   string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
    

    // 上传标识
    
    UploadFlag   string `json:"upload_flag,omitempty" xml:"upload_flag,omitempty"`
    

    // 文件名称
    
    UploadFileName   string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
    

    // 处理标识
    
    ProcessFlag   string `json:"process_flag,omitempty" xml:"process_flag,omitempty"`
    

    // 处理日期
    
    ProcessDate   string `json:"process_date,omitempty" xml:"process_date,omitempty"`
    

    // 单据类型
    
    DataType   string `json:"data_type,omitempty" xml:"data_type,omitempty"`
    

    // 药品类型
    
    DrugType   string `json:"drug_type,omitempty" xml:"drug_type,omitempty"`
    

    // 详情信息
    
    NoteInfo   string `json:"note_info,omitempty" xml:"note_info,omitempty"`
    

    // 操作人编码
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 操作人姓名
    
    CorpName   string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
    

    // 单据日期
    
    StoreInOutDate   string `json:"store_in_out_date,omitempty" xml:"store_in_out_date,omitempty"`
    

    // 上传文件路径
    
    UploadFilePath   string `json:"upload_file_path,omitempty" xml:"upload_file_path,omitempty"`
    

    // 发货单位
    
    FromUserId   string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
    

    // 发货单位
    
    FromRefUserId   string `json:"from_ref_user_id,omitempty" xml:"from_ref_user_id,omitempty"`
    

    // 收货单位
    
    ToUserId   string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
    

    // 收货单位
    
    ToRefUserId   string `json:"to_ref_user_id,omitempty" xml:"to_ref_user_id,omitempty"`
    

    // 是否匹配
    
    OrderIsmatched   string `json:"order_ismatched,omitempty" xml:"order_ismatched,omitempty"`
    

    // 企业ID
    
    RefUserId   string `json:"ref_user_id,omitempty" xml:"ref_user_id,omitempty"`
    

    // 处理状态
    
    ResultType   string `json:"result_type,omitempty" xml:"result_type,omitempty"`
    

}
