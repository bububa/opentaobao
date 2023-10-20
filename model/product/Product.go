package product

import (
	"sync"
)

// Product 结构体
type Product struct {
	// 产品的子图片.目前最多支持4张。fields中设置为product_imgs.id、product_imgs.url、product_imgs.position 等形式就会返回相应的字段
	ProductImgs []ProductImg `json:"product_imgs,omitempty" xml:"product_imgs>product_img,omitempty"`
	// 产品的属性图片.比如说黄色对应的产品图片,绿色对应的产品图片。fields中设置为product_prop_imgs.id、 &lt;br/&gt;product_prop_imgs.props、product_prop_imgs.url、product_prop_imgs.position等形式就会返回相应的字段
	ProductPropImgs []ProductPropImg `json:"product_prop_imgs,omitempty" xml:"product_prop_imgs>product_prop_img,omitempty"`
	// 创建时间.格式:yyyy-mm-dd hh:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 外部产品ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 淘宝标准产品编码
	Tsc string `json:"tsc,omitempty" xml:"tsc,omitempty"`
	// 商品类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 产品的关键属性列表.格式：pid:vid;pid:vid
	Props string `json:"props,omitempty" xml:"props,omitempty"`
	// 产品的关键属性字符串列表.比如:品牌:诺基亚;型号:N73(&lt;strong&gt;注：&lt;/strong&gt;&lt;font color=&#34;red&#34;&gt;属性名称中的冒号&#34;:&#34;被转换为：&#34;#cln#&#34;;  分号&#34;;&#34;被转换为：&#34;#scln#&#34;&lt;/font&gt;)
	PropsStr string `json:"props_str,omitempty" xml:"props_str,omitempty"`
	// 产品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 产品的非关键属性列表.格式:pid:vid;pid:vid.
	Binds string `json:"binds,omitempty" xml:"binds,omitempty"`
	// 产品的非关键属性字符串列表.格式同props_str(&lt;strong&gt;注：&lt;/strong&gt;&lt;font color=&#34;red&#34;&gt;属性名称中的冒号&#34;:&#34;被转换为：&#34;#cln#&#34;;  分号&#34;;&#34;被转换为：&#34;#scln#&#34;&lt;/font&gt;)
	BindsStr string `json:"binds_str,omitempty" xml:"binds_str,omitempty"`
	// 产品的销售属性列表.格式:pid:vid;pid:vid
	SaleProps string `json:"sale_props,omitempty" xml:"sale_props,omitempty"`
	// 产品的销售属性字符串列表.格式同props_str(&lt;strong&gt;注：&lt;/strong&gt;&lt;font color=&#34;red&#34;&gt;属性名称中的冒号&#34;:&#34;被转换为：&#34;#cln#&#34;;  分号&#34;;&#34;被转换为：&#34;#scln#&#34;&lt;/font&gt;)
	SalePropsStr string `json:"sale_props_str,omitempty" xml:"sale_props_str,omitempty"`
	// 产品的市场价.单位为元.精确到2位小数;如:200.07
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 产品的描述.最大25000个字节
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 产品的主图片地址.(绝对地址,格式:http://host/image_path)
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 修改时间.格式:yyyy-mm-dd hh:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。
	PropertyAlias string `json:"property_alias,omitempty" xml:"property_alias,omitempty"`
	// 用户自定义属性,结构：pid1:value1;pid2:value2 例如：“20000:优衣库”，表示“品牌:优衣库”
	CustomerProps string `json:"customer_props,omitempty" xml:"customer_props,omitempty"`
	// 产品卖点描述，长度限制20个汉字
	SellPt string `json:"sell_pt,omitempty" xml:"sell_pt,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 商品类目ID.必须是叶子类目ID
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 当前状态(0 商家确认 1 屏蔽 3 小二确认 2 未确认 -1 删除)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 垂直市场,如：3（3C），4（鞋城）
	VerticalMarket int64 `json:"vertical_market,omitempty" xml:"vertical_market,omitempty"`
}

var poolProduct = sync.Pool{
	New: func() any {
		return new(Product)
	},
}

// GetProduct() 从对象池中获取Product
func GetProduct() *Product {
	return poolProduct.Get().(*Product)
}

// ReleaseProduct 释放Product
func ReleaseProduct(v *Product) {
	v.ProductImgs = v.ProductImgs[:0]
	v.ProductPropImgs = v.ProductPropImgs[:0]
	v.Created = ""
	v.OuterId = ""
	v.Tsc = ""
	v.CatName = ""
	v.Props = ""
	v.PropsStr = ""
	v.Name = ""
	v.Binds = ""
	v.BindsStr = ""
	v.SaleProps = ""
	v.SalePropsStr = ""
	v.Price = ""
	v.Desc = ""
	v.PicUrl = ""
	v.Modified = ""
	v.PropertyAlias = ""
	v.CustomerProps = ""
	v.SellPt = ""
	v.ProductId = 0
	v.Cid = 0
	v.Status = 0
	v.VerticalMarket = 0
	poolProduct.Put(v)
}
