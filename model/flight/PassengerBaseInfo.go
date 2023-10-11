package flight

// PassengerBaseInfo 结构体
type PassengerBaseInfo struct {
	// 身份证前4位，表明省市，仅当card_type=0时，该属性有值
	IdCard string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 年龄，可能为空
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 0:身份证,1:护照,2:学生证,3:军官证,4:回乡证,5:台胞证,6:港澳通行证,7:国际海员,8:外国人永久居留证,9:其它证件,10:警官证,11:士兵证,12:台湾通行证,13:入台证,14:户口簿,15:出生证明,16:驾驶证,17:港澳居民居住证,18:台湾居民居住证
	CardType int64 `json:"card_type,omitempty" xml:"card_type,omitempty"`
}
