package baichuan

// PasswordRuleResultDto 
/* model for simplify = false
type PasswordRuleResultDto struct {

    // level
    
    Level   string `json:"level,omitempty"`
    

    // miaoPasswordRegular
    
    MiaoPasswordRegulars  struct {
        String  []string `json:"string,omitempty"`
    } `json:"miao_password_regulars,omitempty"`
    

    // passwordRegular
    
    PasswordRegulars  struct {
        String  []string `json:"string,omitempty"`
    } `json:"password_regulars,omitempty"`
    

}
*/

// PasswordRuleResultDto 
type PasswordRuleResultDto struct {

    // level
    Level   string `json:"level,omitempty"`

    // miaoPasswordRegular
    MiaoPasswordRegulars   []string `json:"miao_password_regulars,omitempty"`

    // passwordRegular
    PasswordRegulars   []string `json:"password_regulars,omitempty"`

}
