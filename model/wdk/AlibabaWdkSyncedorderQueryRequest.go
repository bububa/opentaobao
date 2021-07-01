package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口查询同步订单 API请求
alibaba.wdk.syncedorder.query

外部商户查询同步到五道口的订单
*/
type AlibabaWdkSyncedorderQueryAPIRequest struct {
    model.Params
    // 门店ID
    _storeId   string
    // 序列号
    _serialNum   string
}

// 初始化AlibabaWdkSyncedorderQueryAPIRequest对象
func NewAlibabaWdkSyncedorderQueryRequest() *AlibabaWdkSyncedorderQueryAPIRequest{
    return &AlibabaWdkSyncedorderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.syncedorder.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkSyncedorderQueryAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetStoreId() string {
    return r._storeId
}
// SerialNum Setter
// 序列号
func (r *AlibabaWdkSyncedorderQueryAPIRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetSerialNum() string {
    return r._serialNum
}
