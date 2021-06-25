package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品ID解密 APIResponse
alibaba.icbu.product.id.decrypt

对混淆的产品ID解密
*/
type AlibabaIcbuProductIdDecryptAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuProductIdDecryptResponse `json:"alibaba_icbu_product_id_decrypt_response,omitempty"`
}

type AlibabaIcbuProductIdDecryptResponse struct {

    // 商品ID
    Id   int64 `json:"id,omitempty"`

}
