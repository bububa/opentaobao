package alihealthoutflow

// PrescriptionGetByVerifyRequest 
type PrescriptionGetByVerifyRequest struct {

    // 药店编码（必选）
    
    DrugStoreCode   string `json:"drug_store_code,omitempty" xml:"drug_store_code,omitempty"`
    

    // 核销码（必选）
    
    VerifyCode   string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
    

    // 药店名称（必选）
    
    DrugStoreName   string `json:"drug_store_name,omitempty" xml:"drug_store_name,omitempty"`
    

    // 药师姓名（必选）
    
    PharmacistName   string `json:"pharmacist_name,omitempty" xml:"pharmacist_name,omitempty"`
    

}
