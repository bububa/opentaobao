package jipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】改签通知单列表 API返回值 
taobao.alitrip.supplier.modify.list

提供供应商查询改签通知单列表
*/
type TaobaoAlitripSupplierModifyListAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripSupplierModifyListResponse
}

// 【机票代理商订单】改签通知单列表 成功返回结果
type TaobaoAlitripSupplierModifyListResponse struct {
    XMLName xml.Name `xml:"alitrip_supplier_modify_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 改签信息列表
    OrderList   []BbSyncOrderDTO `json:"order_list,omitempty" xml:"order_list>bb_sync_order_dto,omitempty"`
}
