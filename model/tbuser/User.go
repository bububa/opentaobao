package tbuser

import (
	"sync"
)

// User 结构体
type User struct {
	// 用户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 用户头像地址
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 性别。可选值:m(男),f(女)
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 对应账号的OpenUID
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 用户类型。可选值:B(B商家),C(C商家)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 是否受限制。可选值:limited(受限制),unlimited(不受限)
	AutoRepost string `json:"auto_repost,omitempty" xml:"auto_repost,omitempty"`
	// 有无实名认证。可选值:authentication(实名认证),not authentication(没有认证)
	PromotedType string `json:"promoted_type,omitempty" xml:"promoted_type,omitempty"`
	// 状态。可选值:normal(正常),inactive(未激活),delete(删除),reeze(冻结),supervise(监管)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 有无绑定。可选值:bind(绑定),notbind(未绑定)
	AlipayBind string `json:"alipay_bind,omitempty" xml:"alipay_bind,omitempty"`
	// 用户的全站vip信息，可取值如下：c(普通会员),asso_vip(荣誉会员)，vip1,vip2,vip3,vip4,vip5,vip6(六个等级的正式vip会员)，共8种取值，其中asso_vip是由vip会员衰退而成，与主站上的vip0对应。
	VipInfo string `json:"vip_info,omitempty" xml:"vip_info,omitempty"`
	// 用户参与垂直市场类型。shoes表示鞋城垂直市场用户，3C表示3C垂直市场用户。多个类型之间用&amp;quot;,&amp;quot;分隔。如：一个用户既是3C用户又是鞋城用户，那么这个字段返回就是&amp;quot;3C,shoes&amp;quot;。如果用户不是垂直市场用户，此字段返回为&amp;quot;&amp;quot;。
	VerticalMarket string `json:"vertical_market,omitempty" xml:"vertical_market,omitempty"`
	// 用户数字ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 卖家信用
	SellerCredit *UserCredit `json:"seller_credit,omitempty" xml:"seller_credit,omitempty"`
	// 可上传商品图片数量
	ItemImgNum int64 `json:"item_img_num,omitempty" xml:"item_img_num,omitempty"`
	// 单张商品图片最大容量(商品主图大小)。单位:k
	ItemImgSize int64 `json:"item_img_size,omitempty" xml:"item_img_size,omitempty"`
	// 可上传属性图片数量
	PropImgNum int64 `json:"prop_img_num,omitempty" xml:"prop_img_num,omitempty"`
	// 单张销售属性图片最大容量（非主图的商品图片和商品属性图片）。单位:k
	PropImgSize int64 `json:"prop_img_size,omitempty" xml:"prop_img_size,omitempty"`
	// 是否购买多图服务。可选值:true(是),false(否)
	HasMorePic bool `json:"has_more_pic,omitempty" xml:"has_more_pic,omitempty"`
	// 是否参加消保
	ConsumerProtection bool `json:"consumer_protection,omitempty" xml:"consumer_protection,omitempty"`
	// 是否订阅了淘宝天下杂志
	MagazineSubscribe bool `json:"magazine_subscribe,omitempty" xml:"magazine_subscribe,omitempty"`
	// 用户是否为网游用户，属于隐私信息，需要登陆才能查看自己的。 目前仅供taobao.user.get使用
	OnlineGaming bool `json:"online_gaming,omitempty" xml:"online_gaming,omitempty"`
	// 是否是无名良品用户，true or false
	Liangpin bool `json:"liangpin,omitempty" xml:"liangpin,omitempty"`
	// 卖家是否签署食品卖家承诺协议
	SignFoodSellerPromise bool `json:"sign_food_seller_promise,omitempty" xml:"sign_food_seller_promise,omitempty"`
	// 用户作为卖家是否开过店
	HasShop bool `json:"has_shop,omitempty" xml:"has_shop,omitempty"`
	// 是否24小时闪电发货(实物类)
	IsLightningConsignment bool `json:"is_lightning_consignment,omitempty" xml:"is_lightning_consignment,omitempty"`
	// 表示用户是否具备修改商品减库存逻辑的权限（一共有拍下减库存和付款减库存两种逻辑） 值含义： 1）true：是 2）false：否。
	HasSubStock bool `json:"has_sub_stock,omitempty" xml:"has_sub_stock,omitempty"`
	// 用户是否是金牌卖家
	IsGoldenSeller bool `json:"is_golden_seller,omitempty" xml:"is_golden_seller,omitempty"`
	// 是否是特价版商家，需要field查询
	IsTjbSeller bool `json:"is_tjb_seller,omitempty" xml:"is_tjb_seller,omitempty"`
}

var poolUser = sync.Pool{
	New: func() any {
		return new(User)
	},
}

// GetUser() 从对象池中获取User
func GetUser() *User {
	return poolUser.Get().(*User)
}

// ReleaseUser 释放User
func ReleaseUser(v *User) {
	v.Nick = ""
	v.Avatar = ""
	v.Sex = ""
	v.OpenUid = ""
	v.Type = ""
	v.AutoRepost = ""
	v.PromotedType = ""
	v.Status = ""
	v.AlipayBind = ""
	v.VipInfo = ""
	v.VerticalMarket = ""
	v.UserId = 0
	v.SellerCredit = nil
	v.ItemImgNum = 0
	v.ItemImgSize = 0
	v.PropImgNum = 0
	v.PropImgSize = 0
	v.HasMorePic = false
	v.ConsumerProtection = false
	v.MagazineSubscribe = false
	v.OnlineGaming = false
	v.Liangpin = false
	v.SignFoodSellerPromise = false
	v.HasShop = false
	v.IsLightningConsignment = false
	v.HasSubStock = false
	v.IsGoldenSeller = false
	v.IsTjbSeller = false
	poolUser.Put(v)
}
