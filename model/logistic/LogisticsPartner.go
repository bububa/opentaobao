package logistic

// LogisticsPartner 
/* model for simplify = false
type LogisticsPartner struct {

    // 揽收说明信息
    
    CoverRemark   string `json:"cover_remark,omitempty"`
    

    // 不可送达的说明信息
    
    UncoverRemark   string `json:"uncover_remark,omitempty"`
    

    // 物流公司揽收和资费详细信息
    
    Carriage  *struct {
        CarriageDetail  *CarriageDetail `json:"carriage_detail,omitempty"`
    } `json:"carriage,omitempty"`
    

    // 物流公司详细信息
    
    Partner  *struct {
        PartnerDetail  *PartnerDetail `json:"partner_detail,omitempty"`
    } `json:"partner,omitempty"`
    

}
*/

// LogisticsPartner 
type LogisticsPartner struct {

    // 揽收说明信息
    CoverRemark   string `json:"cover_remark,omitempty"`

    // 不可送达的说明信息
    UncoverRemark   string `json:"uncover_remark,omitempty"`

    // 物流公司揽收和资费详细信息
    Carriage   *CarriageDetail `json:"carriage,omitempty"`

    // 物流公司详细信息
    Partner   *PartnerDetail `json:"partner,omitempty"`

}
