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
type AlibabaWdkSyncedorderQueryRequest struct {
    model.Params
    // 门店ID
    _storeId   string
    // 序列号
    _serialNum   string
}

// 初始化AlibabaWdkSyncedorderQueryRequest对象
func NewAlibabaWdkSyncedorderQueryRequest() *AlibabaWdkSyncedorderQueryRequest{
    return &AlibabaWdkSyncedorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSyncedorderQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.syncedorder.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSyncedorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkSyncedorderQueryRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkSyncedorderQueryRequest) GetStoreId() string {
    return r._storeId
}
// SerialNum Setter
// 序列号
func (r *AlibabaWdkSyncedorderQueryRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r AlibabaWdkSyncedorderQueryRequest) GetSerialNum() string {
    return r._serialNum
}
