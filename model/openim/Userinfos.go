package openim

// Userinfos 结构体
type Userinfos struct {
	// 昵称，最大长度64字节
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 头像url，最大长度512字节
	IconUrl string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// email地址，最大长度128字节
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 手机号码，最大长度16字节
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 淘宝账号，最大长度64字节
	Taobaoid string `json:"taobaoid,omitempty" xml:"taobaoid,omitempty"`
	// im用户名，最大长度64字节
	Userid string `json:"userid,omitempty" xml:"userid,omitempty"`
	// im密码，最大长度64字节
	Password string `json:"password,omitempty" xml:"password,omitempty"`
	// remark，最大长度128字节
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展字段（Json），最大长度4096字节
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 职位，最大长度128字节
	Career string `json:"career,omitempty" xml:"career,omitempty"`
	// vip（Json），最大长度512字节
	Vip string `json:"vip,omitempty" xml:"vip,omitempty"`
	// 地址，最大长度512
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 名字，最大长度64
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 年龄
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 性别。M: 男。 F：女
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 微信，最大长度64字节
	Wechat string `json:"wechat,omitempty" xml:"wechat,omitempty"`
	// qq，最大长度20字节
	Qq string `json:"qq,omitempty" xml:"qq,omitempty"`
	// 微博，最大长度256字节
	Weibo string `json:"weibo,omitempty" xml:"weibo,omitempty"`
	// 用户激活状态，0表示未激活，1表示激活
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 最后更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}
