package openmall

// RefundMessagePic 
type RefundMessagePic struct {

    // 退款单图片留言
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 退款单图片地址
    
    Pic   string `json:"pic,omitempty" xml:"pic,omitempty"`
    

    // 使用taobao.openmall.refund.image.upload得到的上传token
    
    PicToken   string `json:"pic_token,omitempty" xml:"pic_token,omitempty"`
    

}
