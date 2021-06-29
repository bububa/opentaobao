package drugtrace

// CodeDrugInfoDTO 
type CodeDrugInfoDTO struct {
    // 追溯码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 码生产信息对象
    CodeProduceInfoDTO   *CodeProduceInfoDto `json:"code_produce_info_d_t_o,omitempty" xml:"code_produce_info_d_t_o,omitempty"`
    // 药品基本信息对象
    DrugEntBaseDTO   *DrugEntBaseDto `json:"drug_ent_base_d_t_o,omitempty" xml:"drug_ent_base_d_t_o,omitempty"`
}
