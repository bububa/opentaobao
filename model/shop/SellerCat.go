package shop

// SellerCat 
type SellerCat struct {
    // 卖家自定义类目编号
    Cid   int64 `json:"cid,omitempty" xml:"cid,omitempty"`
    // 父类目编号，值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目
    ParentCid   int64 `json:"parent_cid,omitempty" xml:"parent_cid,omitempty"`
    // 卖家自定义类目名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 链接图片地址
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 该类目在页面上的排序位置
    SortOrder   int64 `json:"sort_order,omitempty" xml:"sort_order,omitempty"`
    // 店铺类目类型：可选值：manual_type：手动分类，new_type：新品上价， tree_type：二三级类目树 ，property_type：属性叶子类目树， brand_type：品牌推广
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
