package ma

// XCodeTo 结构体
type XCodeTo struct {
	// 二维码图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 最后修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 有效期开始时间
	LifeStart string `json:"life_start,omitempty" xml:"life_start,omitempty"`
	// 短连接
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 短连接关键字符
	ShortName string `json:"short_name,omitempty" xml:"short_name,omitempty"`
	// 有效期结束时间，使用短链接跳转的，结束后不再可以访问
	LifeEnd string `json:"life_end,omitempty" xml:"life_end,omitempty"`
	// 记录ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 码的状态，1-生效，0-未生效，-1-逻辑删除。
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 所属用户ID，如果入参没有用户登录信息，则随机生成
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
