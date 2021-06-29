package alihealthmedical

// OuterDoctorPublishRequest 
type OuterDoctorPublishRequest struct {

    // 三方机构医生id
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 医生姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 0-未填写；1-男；2-女
    
    Sex   int64 `json:"sex,omitempty" xml:"sex,omitempty"`
    

    // 医生简介
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 医生头像
    
    UserPicUrl   string `json:"user_pic_url,omitempty" xml:"user_pic_url,omitempty"`
    

    // 医生寄语
    
    SendWord   string `json:"send_word,omitempty" xml:"send_word,omitempty"`
    

    // 医生手机号
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 证件类型，1-身份证
    
    IdentityType   int64 `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
    

    // 证件号码
    
    IdentityNo   string `json:"identity_no,omitempty" xml:"identity_no,omitempty"`
    

    // 证件图片
    
    IdentityPic   string `json:"identity_pic,omitempty" xml:"identity_pic,omitempty"`
    

    // 医师执业证书编号
    
    QualificationNo   string `json:"qualification_no,omitempty" xml:"qualification_no,omitempty"`
    

    // 医师执业证书发证日期
    
    QualificationIssuingTime   string `json:"qualification_issuing_time,omitempty" xml:"qualification_issuing_time,omitempty"`
    

    // 医师执业证书图片
    
    QualificationPic   string `json:"qualification_pic,omitempty" xml:"qualification_pic,omitempty"`
    

    // 医师资格证书发证日期
    
    QuaCertificateTime   string `json:"qua_certificate_time,omitempty" xml:"qua_certificate_time,omitempty"`
    

    // 医师资格证书图片
    
    QuaCertificatePic   string `json:"qua_certificate_pic,omitempty" xml:"qua_certificate_pic,omitempty"`
    

    // 第一执业医院id
    
    YkHospitalId   string `json:"yk_hospital_id,omitempty" xml:"yk_hospital_id,omitempty"`
    

    // 第一执业医院名称
    
    YkHospitalName   string `json:"yk_hospital_name,omitempty" xml:"yk_hospital_name,omitempty"`
    

    // 医院级别。三甲、三乙、三级未定、二甲、二乙、二级未定、一甲、一乙、一级未定、未定级
    
    HospitalLevelName   string `json:"hospital_level_name,omitempty" xml:"hospital_level_name,omitempty"`
    

    // 执业一级科室id
    
    YkFirstDepartId   string `json:"yk_first_depart_id,omitempty" xml:"yk_first_depart_id,omitempty"`
    

    // 执业一级科室名称
    
    YkFirstDepartName   string `json:"yk_first_depart_name,omitempty" xml:"yk_first_depart_name,omitempty"`
    

    // 执业二级科室id
    
    YkSecondDepartId   string `json:"yk_second_depart_id,omitempty" xml:"yk_second_depart_id,omitempty"`
    

    // 执业二级科室名称
    
    YkSecondDepartName   string `json:"yk_second_depart_name,omitempty" xml:"yk_second_depart_name,omitempty"`
    

    // 医生职称
    
    DoctorTitleName   string `json:"doctor_title_name,omitempty" xml:"doctor_title_name,omitempty"`
    

    // 医生擅长疾病code
    
    DiseaseCode   string `json:"disease_code,omitempty" xml:"disease_code,omitempty"`
    

    // 医师资格证书编号
    
    QuaCertificateNo   string `json:"qua_certificate_no,omitempty" xml:"qua_certificate_no,omitempty"`
    

    // 互联网医院id
    
    HospitalId   string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
    

    // 健康标准二级科室ID
    
    JkSecondStDepartId   int64 `json:"jk_second_st_depart_id,omitempty" xml:"jk_second_st_depart_id,omitempty"`
    

    // 健康标准二级科室名称
    
    JkSecondStDepartName   string `json:"jk_second_st_depart_name,omitempty" xml:"jk_second_st_depart_name,omitempty"`
    

    // 医生uuid.新增医生该字段为空，修改医生该字段必传
    
    DoctorUuid   string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
    

}
