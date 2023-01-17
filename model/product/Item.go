package product

// Item 结构体
type Item struct {
	// 商品图片列表(包括主图)。fields中只设置item_img可以返回ItemImg结构体中所有字段，如果设置为item_img.id、item_img.url、item_img.position等形式就只会返回相应的字段
	ItemImgs []ItemImg `json:"item_imgs,omitempty" xml:"item_imgs>item_img,omitempty"`
	// 商品属性图片列表。fields中只设置prop_img可以返回PropImg结构体中所有字段，如果设置为prop_img.id、prop_img.url、prop_img.properties、prop_img.position等形式就只会返回相应的字段
	PropImgs []PropImg `json:"prop_imgs,omitempty" xml:"prop_imgs>prop_img,omitempty"`
	// Sku列表。fields中只设置sku可以返回Sku结构体中所有字段，如果设置为sku.sku_id、sku.properties、sku.quantity等形式就只会返回相应的字段
	Skus []Sku `json:"skus,omitempty" xml:"skus>sku,omitempty"`
	// 商品视频列表(目前只支持单个视频关联)。fields中只设置video可以返回Video结构体中所有字段，如果设置为video.id、video.video_id、&lt;br/&gt;video.url等形式就只会返回相应的字段
	Videos []Video `json:"videos,omitempty" xml:"videos>video,omitempty"`
	// 商品id(注意：iid近期即将废弃，请用num_iid参数)
	Iid string `json:"iid,omitempty" xml:"iid,omitempty"`
	// 商品url
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// 商品标题,不能超过60字节
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 商品类型(fixed:一口价;auction:拍卖)注：取消团购
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 商品所属的店铺内卖家自定义类目列表
	SellerCids string `json:"seller_cids,omitempty" xml:"seller_cids,omitempty"`
	// 商品属性 格式：pid:vid;pid:vid
	Props string `json:"props,omitempty" xml:"props,omitempty"`
	// 用户自行输入的类目属性ID串。结构：&#34;pid1,pid2,pid3&#34;，如：&#34;20000&#34;（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
	InputPids string `json:"input_pids,omitempty" xml:"input_pids,omitempty"`
	// 用户自行输入的子属性名和属性值，结构:&#34;父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....&#34;,如：“耐克;耐克系列;科比系列;科比系列;2K5”，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过 3999字节。
	InputStr string `json:"input_str,omitempty" xml:"input_str,omitempty"`
	// 商品主图片地址
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 商品新旧程度(全新:new，闲置:unused，二手：second)
	StuffStatus string `json:"stuff_status,omitempty" xml:"stuff_status,omitempty"`
	// 商品上传后的状态。onsale出售中，instock库中
	ApproveStatus string `json:"approve_status,omitempty" xml:"approve_status,omitempty"`
	// 属性值别名
	PropertyAlias string `json:"property_alias,omitempty" xml:"property_alias,omitempty"`
	// 商家外部编码(可与商家外部系统对接)
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 秒杀商品类型。打上秒杀标记的商品，用户只能下架并不能再上架，其他任何编辑或删除操作都不能进行。如果用户想取消秒杀标记，需要联系小二进行操作。如果秒杀结束需要自由编辑请联系活动负责人（小二）去掉秒杀标记。可选类型&lt;br/&gt;web_only(只能通过web网络秒杀)&lt;br/&gt;wap_only(只能通过wap网络秒杀)&lt;br/&gt;web_and_wap(既能通过web秒杀也能通过wap秒杀)
	SecondKill string `json:"second_kill,omitempty" xml:"second_kill,omitempty"`
	// 商品属性名称。标识着props内容里面的pid和vid所对应的名称。格式为：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2&amp;hellip;&amp;hellip;(&lt;strong&gt;注：&lt;/strong&gt;&lt;font color=&#34;red&#34;&gt;属性名称中的冒号&amp;quot;:&amp;quot;被转换为：&amp;quot;#cln#&amp;quot;;  分号&amp;quot;;&amp;quot;被转换为：&amp;quot;#scln#&amp;quot;&lt;/font&gt;)
	PropsName string `json:"props_name,omitempty" xml:"props_name,omitempty"`
	// 商品级别的条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品优惠价格
	PromotedPrice string `json:"promoted_price,omitempty" xml:"promoted_price,omitempty"`
	// 商品3:4比例主图
	MainPic34 string `json:"main_pic34,omitempty" xml:"main_pic34,omitempty"`
	// 商品竖图
	UprightImageUrl string `json:"upright_image_url,omitempty" xml:"upright_image_url,omitempty"`
	// 商品数字id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品所属的叶子类目 id
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 商品价格，格式：5.00；单位：元；精确到：分
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// 宝贝所属产品的id(可能为空). 该字段可以通过taobao.products.search 得到
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 返点比例
	AuctionPoint int64 `json:"auction_point,omitempty" xml:"auction_point,omitempty"`
	// 虚拟商品的状态字段
	IsVirtual bool `json:"is_virtual,omitempty" xml:"is_virtual,omitempty"`
}
