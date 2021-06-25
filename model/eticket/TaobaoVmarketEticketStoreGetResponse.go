package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取电子凭证预约门店信息 APIResponse
taobao.vmarket.eticket.store.get

用于给外部商家查询电子凭证预约门店信息
*/
type TaobaoVmarketEticketStoreGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketStoreGetResponse `json:"taobao_vmarket_eticket_store_get_response,omitempty"`
}

type TaobaoVmarketEticketStoreGetResponse struct {

    // 商户id
    StoreId   int64 `json:"store_id,omitempty"`

    // 商户地址
    Address   string `json:"address,omitempty"`

    // 商户名称
    Name   string `json:"name,omitempty"`

    // 所在城市
    City   string `json:"city,omitempty"`

    // 省份
    Province   string `json:"province,omitempty"`

    // 区
    District   string `json:"district,omitempty"`

    // 联系电话
    Contract   string `json:"contract,omitempty"`

    // 自有卖家导入门店的时候，可以把自己系统门店信息的主键或者唯一key传入，用于快速匹配
    Selfcode   string `json:"selfcode,omitempty"`

}
