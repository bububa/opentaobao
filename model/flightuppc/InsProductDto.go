package flightuppc

// InsProductDto 结构体
type InsProductDto struct {
	// 副标题：同一保险在不同页面可以有不同副标题
	SubTitles string `json:"sub_titles,omitempty" xml:"sub_titles,omitempty"`
	// 气泡
	Bubble string `json:"bubble,omitempty" xml:"bubble,omitempty"`
	// 文案
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 利益点：一个保险可以有多个利益点，如：意外保障¥500万+意外医疗¥3万；同一利益点在不同页面可以有不同描述，如：OTA页、OTA浮层、下单页、弹屏页
	Interests string `json:"interests,omitempty" xml:"interests,omitempty"`
	// 主标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 保险产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 标签：同一保险在不同页面可以有不同标签
	Labels string `json:"labels,omitempty" xml:"labels,omitempty"`
	// 保险产品价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 保险产品唯一标识
	InsurancePremiumId int64 `json:"insurance_premium_id,omitempty" xml:"insurance_premium_id,omitempty"`
}
