package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增门店扩展属性 APIResponse
taobao.place.store.extend.add

新增授权用户的门店扩展属性
*/
type TaobaoPlaceStoreExtendAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPlaceStoreExtendAddResponse `json:"taobao_place_store_extend_add_response,omitempty"`
}

type TaobaoPlaceStoreExtendAddResponse struct {

    // 是否失败
    Failure   bool `json:"failure,omitempty"`

    // 返回结果：true成功；false失败
    Result   bool `json:"result,omitempty"`

    // 个数
    TotalNum   int64 `json:"total_num,omitempty"`

}
