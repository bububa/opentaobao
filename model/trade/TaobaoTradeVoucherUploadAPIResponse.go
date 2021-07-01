package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝交易凭证上传 API返回值 
taobao.trade.voucher.upload

定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
*/
type TaobaoTradeVoucherUploadAPIResponse struct {
    model.CommonResponse
    TaobaoTradeVoucherUploadAPIResponseModel
}

// 淘宝交易凭证上传 成功返回结果
type TaobaoTradeVoucherUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"trade_voucher_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 上传到多媒体平台的文件
    File   *File `json:"file,omitempty" xml:"file,omitempty"`
}
