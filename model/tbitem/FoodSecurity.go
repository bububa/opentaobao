package tbitem

import (
	"sync"
)

// FoodSecurity 结构体
type FoodSecurity struct {
	// 厂家联系方式
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 产品标准号
	DesignCode string `json:"design_code,omitempty" xml:"design_code,omitempty"`
	// 厂名
	Factory string `json:"factory,omitempty" xml:"factory,omitempty"`
	// 厂址
	FactorySite string `json:"factory_site,omitempty" xml:"factory_site,omitempty"`
	// 食品添加剂
	FoodAdditive string `json:"food_additive,omitempty" xml:"food_additive,omitempty"`
	// 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
	HealthProductNo string `json:"health_product_no,omitempty" xml:"health_product_no,omitempty"`
	// 配料表
	Mix string `json:"mix,omitempty" xml:"mix,omitempty"`
	// 保质期
	Period string `json:"period,omitempty" xml:"period,omitempty"`
	// 储藏方法
	PlanStorage string `json:"plan_storage,omitempty" xml:"plan_storage,omitempty"`
	// 生产许可证号
	PrdLicenseNo string `json:"prd_license_no,omitempty" xml:"prd_license_no,omitempty"`
	// 生产结束日期
	ProductDateEnd string `json:"product_date_end,omitempty" xml:"product_date_end,omitempty"`
	// 生产开始日期
	ProductDateStart string `json:"product_date_start,omitempty" xml:"product_date_start,omitempty"`
	// 进货结束日期，要在生产日期之后
	StockDateEnd string `json:"stock_date_end,omitempty" xml:"stock_date_end,omitempty"`
	// 进货开始日期，要在生产日期之后
	StockDateStart string `json:"stock_date_start,omitempty" xml:"stock_date_start,omitempty"`
	// 供货商
	Supplier string `json:"supplier,omitempty" xml:"supplier,omitempty"`
}

var poolFoodSecurity = sync.Pool{
	New: func() any {
		return new(FoodSecurity)
	},
}

// GetFoodSecurity() 从对象池中获取FoodSecurity
func GetFoodSecurity() *FoodSecurity {
	return poolFoodSecurity.Get().(*FoodSecurity)
}

// ReleaseFoodSecurity 释放FoodSecurity
func ReleaseFoodSecurity(v *FoodSecurity) {
	v.Contact = ""
	v.DesignCode = ""
	v.Factory = ""
	v.FactorySite = ""
	v.FoodAdditive = ""
	v.HealthProductNo = ""
	v.Mix = ""
	v.Period = ""
	v.PlanStorage = ""
	v.PrdLicenseNo = ""
	v.ProductDateEnd = ""
	v.ProductDateStart = ""
	v.StockDateEnd = ""
	v.StockDateStart = ""
	v.Supplier = ""
	poolFoodSecurity.Put(v)
}
