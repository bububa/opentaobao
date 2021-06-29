package xhotelonlineorder

// CreditCard 
type CreditCard struct {
    // 信用卡cvv
    CvvCode   string `json:"cvv_code,omitempty" xml:"cvv_code,omitempty"`
    // 信用卡卡号
    CardNumber   string `json:"card_number,omitempty" xml:"card_number,omitempty"`
    // 客人选择的有效期格式: 月/年
    ExpirationDate   string `json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
    // 卡类型,AX,MC,DC,VA
    CardCode   string `json:"card_code,omitempty" xml:"card_code,omitempty"`
    // 持卡人姓名
    CardHolderName   string `json:"card_holder_name,omitempty" xml:"card_holder_name,omitempty"`
}
