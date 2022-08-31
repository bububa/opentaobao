package waybill

// UserInfoDto 结构体
type UserInfoDto struct {
	// 手机号码（手机号和固定电话不能同时为空），长度小于20
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 姓名，长度小于40
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 固定电话（手机号和固定电话不能同时为空），长度小于20
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 开放地址ID
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 菜鸟地址ID，针对电商平台加密订单场景使用，淘系订单使用oaid，非淘使用caid。
	Caid string `json:"caid,omitempty" xml:"caid,omitempty"`
	// 发货地址需要通过&lt;a href=&#34;http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.3OFCPk&amp;treeId=17&amp;articleId=104860&amp;docType=1&#34;&gt;search接口&lt;/a&gt;
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
}
