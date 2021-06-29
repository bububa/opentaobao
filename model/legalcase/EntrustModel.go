package legalcase

// EntrustModel 
type EntrustModel struct {
    // 附件
    AttachmentList   []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
    // 案件ids
    CaseIds   []int64 `json:"case_ids,omitempty" xml:"case_ids>int64,omitempty"`
    // 调解截止时间
    Deadline   string `json:"deadline,omitempty" xml:"deadline,omitempty"`
    // 备注
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 委托编号
    EntrustCode   string `json:"entrust_code,omitempty" xml:"entrust_code,omitempty"`
    // 委托方名称
    EntrustParty   string `json:"entrust_party,omitempty" xml:"entrust_party,omitempty"`
    // 委托人
    EntrustPeople   string `json:"entrust_people,omitempty" xml:"entrust_people,omitempty"`
    // 委托时间
    EntrustTime   string `json:"entrust_time,omitempty" xml:"entrust_time,omitempty"`
    // 委托项
    EntrustTypes   []string `json:"entrust_types,omitempty" xml:"entrust_types>string,omitempty"`
    // lvms委托编号
    LvmsEntrustCode   string `json:"lvms_entrust_code,omitempty" xml:"lvms_entrust_code,omitempty"`
    // 主要负责律师
    MainLawyer   string `json:"main_lawyer,omitempty" xml:"main_lawyer,omitempty"`
    // 策略和建议
    Suggest   string `json:"suggest,omitempty" xml:"suggest,omitempty"`
    // 供应商编号
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 供应商名称
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
}
