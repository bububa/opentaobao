package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码商二维码图片上传 APIResponse
taobao.vmarket.eticket.qrcode.upload

电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
*/
type TaobaoVmarketEticketQrcodeUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVmarketEticketQrcodeUploadResponse `json:"vmarket_eticket_qrcode_upload_response,omitempty"` 
    TaobaoVmarketEticketQrcodeUploadResponse
}

/* model for simplify = false
type TaobaoVmarketEticketQrcodeUploadResponse struct {

    // 1:成功  其它为失败
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 图片文件名称
    
    ImgFilename   string `json:"img_filename,omitempty"`
    

}
*/

type TaobaoVmarketEticketQrcodeUploadResponse struct {

    // 1:成功  其它为失败
    RetCode   int64 `json:"ret_code,omitempty"`

    // 图片文件名称
    ImgFilename   string `json:"img_filename,omitempty"`

}
