package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码商二维码图片上传 APIResponse
taobao.vmarket.eticket.qrcode.upload

电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
*/
type TaobaoVmarketEticketQrcodeUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_qrcode_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1:成功  其它为失败
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"