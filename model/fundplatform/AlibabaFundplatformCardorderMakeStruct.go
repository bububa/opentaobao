package fundplatform

// AlibabaFundplatformCardorderMakeStruct 
type AlibabaFundplatformCardorderMakeStruct struct {

    // 收件人邮编
    
    ReceiverZipCode   string `json:"receiver_zip_code,omitempty" xml:"receiver_zip_code,omitempty"`
    

    // 收件人地址
    
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    

    // 收件人联系方式
    
    ReceiverNumber   string `json:"receiver_number,omitempty" xml:"receiver_number,omitempty"`
    

    // 收件人名称
    
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    

    // 卡模板编号
    
    TemplateNo   string `json:"template_no,omitempty" xml:"template_no,omitempty"`
    

    // 卡外观图片地址
    
    PictureUrl   string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
    

    // 当前卡模板制卡数量
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

    // 卡面额，单元分
    
    ParValue   string `json:"par_value,omitempty" xml:"par_value,omitempty"`
    

    // 该模板生成的卡名称
    
    CardName   string `json:"card_name,omitempty" xml:"card_name,omitempty"`
    

    // 若制卡请求接收成功，返回制卡商内部对应此订单的编号。目前未实际使用，可以随意填写
    
    MakingCardNo   string `json:"making_card_no,omitempty" xml:"making_card_no,omitempty"`
    

    // 错误详情
    
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`
    

    // 错误CODE
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

    // 是否调用成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
