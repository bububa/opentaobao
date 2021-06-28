package tmallhk

// CtsRing 
type CtsRing struct {

    // 成品生产完成时间
    
    CompletedTime   string `json:"completed_time,omitempty" xml:"completed_time,omitempty"`
    

    // 额外信息
    
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
    

    // 戒托货品ID
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 戒托镶嵌完成时间
    
    MountTime   string `json:"mount_time,omitempty" xml:"mount_time,omitempty"`
    

    // 戒托订单号
    
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    

    // 戒托商品Id
    
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 戒托生产完成时间
    
    RingTime   string `json:"ring_time,omitempty" xml:"ring_time,omitempty"`
    

    // 戒托子订单号
    
    SubOrderNo   string `json:"sub_order_no,omitempty" xml:"sub_order_no,omitempty"`
    

}
