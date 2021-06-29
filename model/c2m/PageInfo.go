package c2m

// PageInfo 
type PageInfo struct {
    // 总条数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 总页数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 页大小，固定50
    Pages   int64 `json:"pages,omitempty" xml:"pages,omitempty"`
    // 第几页
    PageNum   int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
    // 订单信息
    List   []CompanyOrderInfoVo `json:"list,omitempty" xml:"list>company_order_info_vo,omitempty"`
}
