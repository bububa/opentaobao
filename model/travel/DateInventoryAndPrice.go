package travel

// DateInventoryAndPrice 
type DateInventoryAndPrice struct {
    // 销售日期
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // Sku的销售价格。单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
    Stock   int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}
