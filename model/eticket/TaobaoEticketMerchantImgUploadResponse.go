package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码商上传二维码图片 APIResponse
taobao.eticket.merchant.img.upload

电子凭证的码商可以通过这个接口，上传二维码图片
*/
type TaobaoEticketMerchantImgUploadAPIResponse struct {
    model.CommonResponse
    Response *TaobaoEticketMerchantImgUploadResponse `json:"taobao_eticket_merchant_img_upload_response,omitempty"`
}

type TaobaoEticketMerchantImgUploadResponse struct {

    // 回复对象
    RespBody   *UploadImgCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
