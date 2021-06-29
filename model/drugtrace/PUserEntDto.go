package drugtrace

// PUserEntDto 
type PUserEntDto struct {

    // 机构编码
    
    OrgCode   string `json:"org_code,omitempty" xml:"org_code,omitempty"`
    

    // 原企业名称
    
    EntName   string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
    

    // 新企业名称
    
    EntNameNew   string `json:"ent_name_new,omitempty" xml:"ent_name_new,omitempty"`
    

    // 企业refid
    
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    

    // 企业id
    
    EntId   string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
    

}
