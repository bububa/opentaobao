package ieagency

// IeChangePassengerVo 
type IeChangePassengerVo struct {

    // 出生日期
    
    BirthDate   string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
    

    // 证件过期日期
    
    DocumentExpireDate   string `json:"document_expire_date,omitempty" xml:"document_expire_date,omitempty"`
    

    // 持有人国际
    
    DocumentHolderNationality   string `json:"document_holder_nationality,omitempty" xml:"document_holder_nationality,omitempty"`
    

    // 证件号
    
    DocumentID   string `json:"document_i_d,omitempty" xml:"document_i_d,omitempty"`
    

    // 发证国际
    
    DocumentIssueCountry   string `json:"document_issue_country,omitempty" xml:"document_issue_country,omitempty"`
    

    // 证件类型
    
    DocumentType   int64 `json:"document_type,omitempty" xml:"document_type,omitempty"`
    

    // 性别
    
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 乘机人编号
    
    PassengerId   int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
    

    // 乘机人类型 0-成人,1-儿童,2-留学生
    
    PassengerType   int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
    

}
