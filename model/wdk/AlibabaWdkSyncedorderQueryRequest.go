package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口查询同步订单 APIRequest
alibaba.wdk.syncedorder.query

外部商户查询同步到五道口的订单
*/
type AlibabaWdkSyncedorderQueryRequest struct {
    model.Params

    // 门店ID
    storeId   string 

    // 序列号
    serialNum   string 

}

func NewAlibabaWdkSyncedorderQueryRequest() *AlibabaWdkSyncedorderQueryRequest{
    return &AlibabaWdkSyncedorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSyncedorderQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.syncedorder.query"
}

func (r AlibabaWdkSyncedorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSyncedorderQueryRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkSyncedorderQueryRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkSyncedorderQueryRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r AlibabaWdkSyncedorderQueryRequest) GetSerialNum() string {
    return r.serialNum
}

