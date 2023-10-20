package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSyncedorderQueryAPIRequest 五道口查询同步订单 API请求
// alibaba.wdk.syncedorder.query
//
// 外部商户查询同步到五道口的订单
type AlibabaWdkSyncedorderQueryAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 序列号
	_serialNum string
}

// NewAlibabaWdkSyncedorderQueryRequest 初始化AlibabaWdkSyncedorderQueryAPIRequest对象
func NewAlibabaWdkSyncedorderQueryRequest() *AlibabaWdkSyncedorderQueryAPIRequest {
	return &AlibabaWdkSyncedorderQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSyncedorderQueryAPIRequest) Reset() {
	r._storeId = ""
	r._serialNum = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.syncedorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabaWdkSyncedorderQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSerialNum is SerialNum Setter
// 序列号
func (r *AlibabaWdkSyncedorderQueryAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// GetSerialNum SerialNum Getter
func (r AlibabaWdkSyncedorderQueryAPIRequest) GetSerialNum() string {
	return r._serialNum
}

var poolAlibabaWdkSyncedorderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSyncedorderQueryRequest()
	},
}

// GetAlibabaWdkSyncedorderQueryRequest 从 sync.Pool 获取 AlibabaWdkSyncedorderQueryAPIRequest
func GetAlibabaWdkSyncedorderQueryAPIRequest() *AlibabaWdkSyncedorderQueryAPIRequest {
	return poolAlibabaWdkSyncedorderQueryAPIRequest.Get().(*AlibabaWdkSyncedorderQueryAPIRequest)
}

// ReleaseAlibabaWdkSyncedorderQueryAPIRequest 将 AlibabaWdkSyncedorderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSyncedorderQueryAPIRequest(v *AlibabaWdkSyncedorderQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSyncedorderQueryAPIRequest.Put(v)
}
