package util

// AssetQrCodeDto 
/* model for simplify = false
type AssetQrCodeDto struct {

    // 资产类型
    
    AssetType   string `json:"asset_type,omitempty"`
    

    // 实物来源
    
    EntitySource   string `json:"entity_source,omitempty"`
    

    // 请求生产条码数量
    
    Quantity   string `json:"quantity,omitempty"`
    

    // 配件厂商code(请求参数)
    
    SpareBrandCode   string `json:"spare_brand_code,omitempty"`
    

    // 整机厂商code(请求参数)
    
    DeviceBrandCode   string `json:"device_brand_code,omitempty"`
    

    // 实物SN(请求参数)
    
    Sn   string `json:"sn,omitempty"`
    

    // 阿里侧部件型号(请求参数)
    
    Mpn   string `json:"mpn,omitempty"`
    

    // 生产二维码信息字符串
    
    QrCodeStringList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"qr_code_string_list,omitempty"`
    

    // 配件类型code(请求参数)
    
    SpareCategoryCode   string `json:"spare_category_code,omitempty"`
    

    // 厂商代码(请求参数)
    
    Manufacture   string `json:"manufacture,omitempty"`
    

    // 单批次请求唯一标识
    
    Nonce   string `json:"nonce,omitempty"`
    

}
*/

// AssetQrCodeDto 
type AssetQrCodeDto struct {

    // 资产类型
    AssetType   string `json:"asset_type,omitempty"`

    // 实物来源
    EntitySource   string `json:"entity_source,omitempty"`

    // 请求生产条码数量
    Quantity   string `json:"quantity,omitempty"`

    // 配件厂商code(请求参数)
    SpareBrandCode   string `json:"spare_brand_code,omitempty"`

    // 整机厂商code(请求参数)
    DeviceBrandCode   string `json:"device_brand_code,omitempty"`

    // 实物SN(请求参数)
    Sn   string `json:"sn,omitempty"`

    // 阿里侧部件型号(请求参数)
    Mpn   string `json:"mpn,omitempty"`

    // 生产二维码信息字符串
    QrCodeStringList   []string `json:"qr_code_string_list,omitempty"`

    // 配件类型code(请求参数)
    SpareCategoryCode   string `json:"spare_category_code,omitempty"`

    // 厂商代码(请求参数)
    Manufacture   string `json:"manufacture,omitempty"`

    // 单批次请求唯一标识
    Nonce   string `json:"nonce,omitempty"`

}
