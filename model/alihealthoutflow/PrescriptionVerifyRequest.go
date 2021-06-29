package alihealthoutflow

// PrescriptionVerifyRequest 
type PrescriptionVerifyRequest struct {

    // 核销码（必选）
    
    VerifyCode   string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
    

    // 处方号（必选）
    
    RxNo   string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
    

    // 药师名字（必选）
    
    PharmacistName   string `json:"pharmacist_name,omitempty" xml:"pharmacist_name,omitempty"`
    

    // 药店id（必选）
    
    DrugStoreCode   string `json:"drug_store_code,omitempty" xml:"drug_store_code,omitempty"`
    

    // 药店名字（必选）
    
    DrugStoreName   string `json:"drug_store_name,omitempty" xml:"drug_store_name,omitempty"`
    

    // 是否核销（必选）
    
    IfConfirm   bool `json:"if_confirm,omitempty" xml:"if_confirm,omitempty"`
    

    // 结算id（必选）
    
    OutSettlement   string `json:"out_settlement,omitempty" xml:"out_settlement,omitempty"`
    

    // 药品信息（必选）
    
    DrugList   []DrugDTO `json:"drug_list,omitempty" xml:"drug_list,omitempty"`
    

    // 核销金额（必选）
    
    VerificationAmount   string `json:"verification_amount,omitempty" xml:"verification_amount,omitempty"`
    

}
