package xhotelitem

// AddOwnerParam 结构体
type AddOwnerParam struct {
	// 房东Id，供货商自己库中的房东Id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 房东昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 房东头像地址
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 性别，M-男性,F-女性，U-未知
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 房东的生日（年-月-日 00:00:00）
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 职业
	Profession string `json:"profession,omitempty" xml:"profession,omitempty"`
	// 爱好，用英文逗号分割 如"游泳,爬山"
	Hobbies string `json:"hobbies,omitempty" xml:"hobbies,omitempty"`
	// 房东介绍，长度限制2048个字符
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 房东故事，json串，可以传图片和文本： {"pics":[""],"content":""}
	Story string `json:"story,omitempty" xml:"story,omitempty"`
	// 固定电话。移动电话号码与固定电话号码二者必须填一个
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 移动电话号码。移动电话号码与固定电话号码二者必须填一个
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 房东的真实联系方式，不能填第三方转接号码
	RealContact string `json:"real_contact,omitempty" xml:"real_contact,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 国内固定传"China"；国外必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 经营资质执照类型对应的名称，执照类型1（身份证）：身份证姓名；执照类型2（护照）：护照姓名；执照类型3（营业执照）：企业名称
	LicenseName string `json:"license_name,omitempty" xml:"license_name,omitempty"`
	// 经营资质执照类型对应的编码，执照类型1（身份证）：身份证号；执照类型2（护照）：护照号；执照类型3（营业执照）：营业执照编号
	LicenseNo string `json:"license_no,omitempty" xml:"license_no,omitempty"`
	// 实名认证姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 身份证号
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 房东标签，英文逗号分割，如"超赞,有爱"
	Labels string `json:"labels,omitempty" xml:"labels,omitempty"`
	// 房东类型，枚举 1.个人房东；2.商户经营；3.其他
	OwnerType int64 `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	// 学历 1:小学,2:初中,3:高中,4:本科,5:硕士,6:博士,7:博士后,0:其他
	Qualification int64 `json:"qualification,omitempty" xml:"qualification,omitempty"`
	// 星座 -1:错误，未知, 0:白羊,1:金牛,2:双子,3:巨蟹,4:狮子,5:处女,6:天秤,7:天蝎,8:射手,9:摩羯,10:水瓶,11:双鱼
	Constellation int64 `json:"constellation,omitempty" xml:"constellation,omitempty"`
	// 房东血型，0:未知,1:A型,2:B型,3:AB型,4:O型
	BloodType int64 `json:"blood_type,omitempty" xml:"blood_type,omitempty"`
	// 省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 房东所在城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 经营资质执照类型 ，1 身份证(默认)、2 护照、3 营业执照
	LicenseType int64 `json:"license_type,omitempty" xml:"license_type,omitempty"`
	// 实名验证方式，认证情况:1:身份验证,2:头像验证,4:手机验证,8:邮箱验证,可以二进制叠加,二进制各位代表含义
	Validate int64 `json:"validate,omitempty" xml:"validate,omitempty"`
	// 房东等级评分，0-100，房东等级越高越优秀
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 订单平均确认时长（单位分钟，要求大于0）
	AvgConfirmTime int64 `json:"avg_confirm_time,omitempty" xml:"avg_confirm_time,omitempty"`
	// 订单接单率 0-100，百分比分子
	ConfirmRate int64 `json:"confirm_rate,omitempty" xml:"confirm_rate,omitempty"`
	// 回复率，0-100，百分比分子，数字越大回复率越高
	ResponseRate int64 `json:"response_rate,omitempty" xml:"response_rate,omitempty"`
	// 好评率，0-100，百分比分子，数字越大好评率越高
	PositiveFeedback int64 `json:"positive_feedback,omitempty" xml:"positive_feedback,omitempty"`
}
