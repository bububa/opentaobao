package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一般进口发货 APIRequest
taobao.wlb.imports.general.consign

将订单信息发送到菜鸟海外转运仓；
业务规则：
1）交易订单为待发货状态。
2）单笔订单多个商品，交易金额不能大于1000人民币。
*/
type TaobaoWlbImportsGeneralConsignRequest struct {
    model.Params

    // 交易订单id
    tradeOrderId   int64 

    // 物流资源ID
    resourceId   int64 

    // 仓库编码
    storeCode   string 

    // 第1段物流承运商
    firstLogistics   string 

    // 第1段物流运单号
    firstWaybillno   string 

    // 卖家发货地址库ID；不填的话取默认发货地址；如果填写的senderId对应多个地址，取第一个
    senderId   int64 

    // 卖家退货地址库ID；不填写的话取默认退货地址；如果填写的cancelId对应多个地址，取第一个
    cancelId   int64 

    // 增值服务编码.多个以逗号分隔
    vasCode   string 

}

func NewTaobaoWlbImportsGeneralConsignRequest() *TaobaoWlbImportsGeneralConsignRequest{
    return &TaobaoWlbImportsGeneralConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.general.consign"
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportsGeneralConsignRequest) SetTradeOrderId(tradeOrderId int64) error {
    r.tradeOrderId = tradeOrderId
    r.Set("trade_order_id", tradeOrderId)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetTradeOrderId() int64 {
    return r.tradeOrderId
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetResourceId(resourceId int64) error {
    r.resourceId = resourceId
    r.Set("resource_id", resourceId)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetResourceId() int64 {
    return r.resourceId
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetFirstLogistics(firstLogistics string) error {
    r.firstLogistics = firstLogistics
    r.Set("first_logistics", firstLogistics)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetFirstLogistics() string {
    return r.firstLogistics
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetFirstWaybillno(firstWaybillno string) error {
    r.firstWaybillno = firstWaybillno
    r.Set("first_waybillno", firstWaybillno)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetFirstWaybillno() string {
    return r.firstWaybillno
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetSenderId(senderId int64) error {
    r.senderId = senderId
    r.Set("sender_id", senderId)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetSenderId() int64 {
    return r.senderId
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetCancelId(cancelId int64) error {
    r.cancelId = cancelId
    r.Set("cancel_id", cancelId)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetCancelId() int64 {
    return r.cancelId
}

func (r *TaobaoWlbImportsGeneralConsignRequest) SetVasCode(vasCode string) error {
    r.vasCode = vasCode
    r.Set("vas_code", vasCode)
    return nil
}

func (r TaobaoWlbImportsGeneralConsignRequest) GetVasCode() string {
    return r.vasCode
}

