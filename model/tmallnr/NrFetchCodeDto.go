package tmallnr

// NrFetchCodeDto 
type NrFetchCodeDto struct {

    // 主订单号
    
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    

    // 发货公司名称
    
    ConsignCompanyName   string `json:"consign_company_name,omitempty" xml:"consign_company_name,omitempty"`
    

    // 业务标识（fn/cn）
    
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

    // 发货公司编码
    
    ConsignCompanyCode   string `json:"consign_company_code,omitempty" xml:"consign_company_code,omitempty"`
    

    // 打印内容
    
    Printdata   string `json:"printdata,omitempty" xml:"printdata,omitempty"`
    

    // 对货码
    
    ShortId   int64 `json:"short_id,omitempty" xml:"short_id,omitempty"`
    

    // 取件码
    
    FetchCode   string `json:"fetch_code,omitempty" xml:"fetch_code,omitempty"`
    

    // 面单号
    
    FaceSheetId   string `json:"face_sheet_id,omitempty" xml:"face_sheet_id,omitempty"`
    

    // 菜鸟生成的标签号
    
    TagNo   string `json:"tag_no,omitempty" xml:"tag_no,omitempty"`
    

    // 物流cp内部的ID号
    
    CpOutId   string `json:"cp_out_id,omitempty" xml:"cp_out_id,omitempty"`
    

    // 收货人的手机号
    
    ReceivePhone   string `json:"receive_phone,omitempty" xml:"receive_phone,omitempty"`
    

    // 收货地址
    
    ReceiveAddr   string `json:"receive_addr,omitempty" xml:"receive_addr,omitempty"`
    

    // 收货人名称
    
    ReceiveName   string `json:"receive_name,omitempty" xml:"receive_name,omitempty"`
    

    // 发货省份
    
    SendProvince   string `json:"send_province,omitempty" xml:"send_province,omitempty"`
    

    // 发货城市
    
    SendCity   string `json:"send_city,omitempty" xml:"send_city,omitempty"`
    

    // 发货详细地址
    
    SendAddr   string `json:"send_addr,omitempty" xml:"send_addr,omitempty"`
    

    // 核销码
    
    WriteOffCode   string `json:"write_off_code,omitempty" xml:"write_off_code,omitempty"`
    

}
