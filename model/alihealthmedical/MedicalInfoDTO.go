package alihealthmedical

// MedicalInfoDto 
type MedicalInfoDto struct {
    // 既往史详情
    PastHistoryDetails   []string `json:"past_history_details,omitempty" xml:"past_history_details>string,omitempty"`
    // 城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 疾病照片、检查单
    DiseasePictures   []string `json:"disease_pictures,omitempty" xml:"disease_pictures>string,omitempty"`
    // 过敏史详情
    AllergiesDetails   []string `json:"allergies_details,omitempty" xml:"allergies_details>string,omitempty"`
    // 主诉
    ChiefComplaint   string `json:"chief_complaint,omitempty" xml:"chief_complaint,omitempty"`
    // 肝功能异常详情
    AbnormalLiverFunctionDetail   string `json:"abnormal_liver_function_detail,omitempty" xml:"abnormal_liver_function_detail,omitempty"`
    // 肾功能异常详情
    AbnormalRenalFunctionDetail   string `json:"abnormal_renal_function_detail,omitempty" xml:"abnormal_renal_function_detail,omitempty"`
    // 是否有过敏史
    HasAllergiesHistory   bool `json:"has_allergies_history,omitempty" xml:"has_allergies_history,omitempty"`
    // 是否（备孕/妊娠/哺乳期）
    IsPregnant   bool `json:"is_pregnant,omitempty" xml:"is_pregnant,omitempty"`
    // 现病史
    PresentIllnessHistory   string `json:"present_illness_history,omitempty" xml:"present_illness_history,omitempty"`
    // 家族史详情
    FamilyHistoryDetails   []string `json:"family_history_details,omitempty" xml:"family_history_details>string,omitempty"`
    // 既往史描述
    PastHistoryDescription   string `json:"past_history_description,omitempty" xml:"past_history_description,omitempty"`
    // 性别
    Sex   string `json:"sex,omitempty" xml:"sex,omitempty"`
    // 是否肝功能异常
    IsLiverFunctionAbnormal   bool `json:"is_liver_function_abnormal,omitempty" xml:"is_liver_function_abnormal,omitempty"`
    // 预产期，如果妊娠则需要填写预产期
    ExpectedConfinementDate   string `json:"expected_confinement_date,omitempty" xml:"expected_confinement_date,omitempty"`
    // 家族史描述
    FamilyHistoryDescription   string `json:"family_history_description,omitempty" xml:"family_history_description,omitempty"`
    // 过敏史描述
    AllergiesDescription   string `json:"allergies_description,omitempty" xml:"allergies_description,omitempty"`
    // 是否肾功能异常
    IsRenalFunctionAbnormal   bool `json:"is_renal_function_abnormal,omitempty" xml:"is_renal_function_abnormal,omitempty"`
    // 已确诊的疾病
    DiagnosedDiseases   []string `json:"diagnosed_diseases,omitempty" xml:"diagnosed_diseases>string,omitempty"`
    // 备孕/妊娠/哺乳
    PregnantType   string `json:"pregnant_type,omitempty" xml:"pregnant_type,omitempty"`
    // 是否有家族史
    HasFamilyHistory   bool `json:"has_family_history,omitempty" xml:"has_family_history,omitempty"`
    // 月龄，单位为“个月”
    Age   string `json:"age,omitempty" xml:"age,omitempty"`
    // 是否有既往史
    HasPastHistory   bool `json:"has_past_history,omitempty" xml:"has_past_history,omitempty"`
}
