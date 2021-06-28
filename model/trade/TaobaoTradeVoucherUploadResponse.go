package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝交易凭证上传 APIResponse
taobao.trade.voucher.upload

定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
*/
type TaobaoTradeVoucherUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTradeVoucherUploadResponse `json:"trade_voucher_upload_response,omitempty"` 
    TaobaoTradeVoucherUploadResponse
}

/* model for simplify = false
type TaobaoTradeVoucherUploadResponse struct {

    // 上传到多媒体平台的文件
    
    File  *struct {
        File  *File `json:"file,omitempty"`
    } `json:"file,omitempty"`
    

}
*/

type TaobaoTradeVoucherUploadResponse struct {

    // 上传到多媒体平台的文件
    File   *File `json:"file,omitempty"`

}
