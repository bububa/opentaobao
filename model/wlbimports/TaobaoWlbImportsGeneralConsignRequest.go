package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一般进口发货 API请求
taobao.wlb.imports.general.consign

将订单信息发送到菜鸟海外转运仓；
业务规则：
1）交易订单为待发货状态。
2）单笔订单多个商品，交易金额不能大于1000人民币。
*/
type TaobaoWlbImportsGeneralConsignAPIRequest struct {
    model.Params
    // 交易订单id
    _tradeOrderId   int64
    // 物流资源ID
    _resourceId   int64
    // 仓库编码
    _storeCode   string
    // 第1段物流承运商
    _firstLogistics   string
    // 第1段物流运单号
    _firstWaybillno   string
    // 卖家发货地址库ID；不填的话取默认发货地址；如果填写的senderId对应多个地址，取第一个
    _senderId   int64
    // 卖家退货地址库ID；不填写的话取默认退货地址；如果填写的cancelId对应多个地址，取第一个
    _cancelId   int64
    // 增值服务编码.多个以逗号分隔
    _vasCode   string
}

// 初始化TaobaoWlbImportsGeneralConsignAPIRequest对象
func NewTaobaoWlbImportsGeneralConsignRequest() *TaobaoWlbImportsGeneralConsignAPIRequest{
    return &TaobaoWlbImportsGeneralConsignAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.general.consign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeOrderId Setter
// 交易订单id
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
    r._tradeOrderId = _tradeOrderId
    r.Set("trade_order_id", _tradeOrderId)
    return nil
}

// TradeOrderId Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetTradeOrderId() int64 {
    return r._tradeOrderId
}
// ResourceId Setter
// 物流资源ID
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetResourceId(_resourceId int64) error {
    r._resourceId = _resourceId
    r.Set("resource_id", _resourceId)
    return nil
}

// ResourceId Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetResourceId() int64 {
    return r._resourceId
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetStoreCode() string {
    return r._storeCode
}
// FirstLogistics Setter
// 第1段物流承运商
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetFirstLogistics(_firstLogistics string) error {
    r._firstLogistics = _firstLogistics
    r.Set("first_logistics", _firstLogistics)
    return nil
}

// FirstLogistics Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetFirstLogistics() string {
    return r._firstLogistics
}
// FirstWaybillno Setter
// 第1段物流运单号
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetFirstWaybillno(_firstWaybillno string) error {
    r._firstWaybillno = _firstWaybillno
    r.Set("first_waybillno", _firstWaybillno)
    return nil
}

// FirstWaybillno Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetFirstWaybillno() string {
    return r._firstWaybillno
}
// SenderId Setter
// 卖家发货地址库ID；不填的话取默认发货地址；如果填写的senderId对应多个地址，取第一个
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetSenderId(_senderId int64) error {
    r._senderId = _senderId
    r.Set("sender_id", _senderId)
    return nil
}

// SenderId Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetSenderId() int64 {
    return r._senderId
}
// CancelId Setter
// 卖家退货地址库ID；不填写的话取默认退货地址；如果填写的cancelId对应多个地址，取第一个
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetCancelId(_cancelId int64) error {
    r._cancelId = _cancelId
    r.Set("cancel_id", _cancelId)
    return nil
}

// CancelId Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetCancelId() int64 {
    return r._cancelId
}
// VasCode Setter
// 增值服务编码.多个以逗号分隔
func (r *TaobaoWlbImportsGeneralConsignAPIRequest) SetVasCode(_vasCode string) error {
    r._vasCode = _vasCode
    r.Set("vas_code", _vasCode)
    return nil
}

// VasCode Getter
func (r TaobaoWlbImportsGeneralConsignAPIRequest) GetVasCode() string {
    return r._vasCode
}
