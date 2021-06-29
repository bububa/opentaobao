package alihealthcrm

// BabyBaseInfoDTO 
type BabyBaseInfoDTO struct {

    // 宝宝id
    
    BabyId   int64 `json:"baby_id,omitempty" xml:"baby_id,omitempty"`
    

    // 所属用户id
    
    TpUserId   int64 `json:"tp_user_id,omitempty" xml:"tp_user_id,omitempty"`
    

    // 头像
    
    HeadPhotoUrl   string `json:"head_photo_url,omitempty" xml:"head_photo_url,omitempty"`
    

    // 预产期
    
    ExpecteDate   string `json:"expecte_date,omitempty" xml:"expecte_date,omitempty"`
    

    // 宝宝状态 0 是育期，1 是孕期
    
    GestationStatus   string `json:"gestation_status,omitempty" xml:"gestation_status,omitempty"`
    

    // 宝宝名称
    
    BabyName   string `json:"baby_name,omitempty" xml:"baby_name,omitempty"`
    

    // 性别
    
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 宝宝名称
    
    Birthday   string `json:"birthday,omitempty" xml:"birthday,omitempty"`
    

    // 提醒标志
    
    RemindFlag   int64 `json:"remind_flag,omitempty" xml:"remind_flag,omitempty"`
    

}
