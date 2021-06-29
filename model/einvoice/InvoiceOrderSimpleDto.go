package einvoice

// InvoiceOrderSimpleDto 
type InvoiceOrderSimpleDto struct {

    // 所绑定的税控设备ID;  入驻成功 & 单机版税控产品时，包含该字段。
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 订购单ID
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 税控产品，产品码由中台定义。
    
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    

    // 服务结束时间，格式yyyy-MM-dd HH:mm:ss 当flow_status=success时必选；
    
    ServEndTime   string `json:"serv_end_time,omitempty" xml:"serv_end_time,omitempty"`
    

    // 服务起始时间，格式yyyy-MM-dd HH:mm:ss  当flow_status=success时必选；
    
    ServStartTime   string `json:"serv_start_time,omitempty" xml:"serv_start_time,omitempty"`
    

}
