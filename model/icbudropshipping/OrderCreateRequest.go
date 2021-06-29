package icbudropshipping

// OrderCreateRequest 
type OrderCreateRequest struct {

    // Provide the order number corresponding to the 3rd party ISV
    
    ChannelReferId   string `json:"channel_refer_id,omitempty" xml:"channel_refer_id,omitempty"`
    

    // Logistics details
    
    LogisticsDetail   *LogisticsDetail `json:"logistics_detail,omitempty" xml:"logistics_detail,omitempty"`
    

    // Payment details
    
    PaymentDetail   *PaymentDetail `json:"payment_detail,omitempty" xml:"payment_detail,omitempty"`
    

    // Product list
    
    ProductList   []TradeEcologyOrderCreateProduct `json:"product_list,omitempty" xml:"product_list,omitempty"`
    

    // Put the order number provided by the 3rd party platform and the name of the 3rd party platform. For example, if the order number is for a transaction made on Shopify, put “Shopify” and the order number. <br />  Platform Names can be case ignored:<br /> Shopify,CommerceHQ,WooCommerce,GrooveKart,BigCommerce
    
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
    

    // order remark
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

}
