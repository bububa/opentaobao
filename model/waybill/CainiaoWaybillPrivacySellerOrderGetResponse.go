package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订单查询 APIResponse
cainiao.waybill.privacy.seller.order.get

商家查询最近100天隐私面单记录
*/
type CainiaoWaybillPrivacySellerOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillPrivacySellerOrderGetResponse `json:"cainiao_waybill_privacy_seller_order_get_response,omitempty"` 
    CainiaoWaybillPrivacySellerOrderGetResponse
}

/* model for simplify = false
type CainiaoWaybillPrivacySellerOrderGetResponse struct {

    // 错误列表
    
    ErrorCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_code_list,omitempty"`
    

    // 是否失败
    
    Failure   bool `json:"failure,omitempty"`
    

    // 第一个错误
    
    OneErrorInfo   string `json:"one_error_info,omitempty"`
    

    // 错误信息
    
    ErrorInfoList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_info_list,omitempty"`
    

    // objectId
    
    ObjectId   string `json:"object_id,omitempty"`
    

    // 返回值
    
    ResponseList  struct {
        CainiaoWaybillPrivacySellerOrderGetModule  []CainiaoWaybillPrivacySellerOrderGetModule `json:"cainiao_waybill_privacy_seller_order_get_module,omitempty"`
    } `json:"response_list,omitempty"`
    

}
*/

type CainiaoWaybillPrivacySellerOrderGetResponse struct {

    // 错误列表
    ErrorCodeList   []string `json:"error_code_list,omitempty"`

    // 是否失败
    Failure   bool `json:"failure,omitempty"`

    // 第一个错误
    OneErrorInfo   string `json:"one_error_info,omitempty"`

    // 错误信息
    ErrorInfoList   []string `json:"error_info_list,omitempty"`

    // objectId
    ObjectId   string `json:"object_id,omitempty"`

    // 返回值
    ResponseList   []CainiaoWaybillPrivacySellerOrderGetModule `json:"response_list,omitempty"`

}
