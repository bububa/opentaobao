package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdksyncedorderqueryAPIRequest 五道口查询同步订单 API请求
// alibaba.wdk.syncedorder.query
//
// 外部商户查询同步到五道口的订单
type AlibabawdksyncedorderqueryAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 序列号
	_serialNum string
}

// NewAlibabawdksyncedorderqueryRequest 初始化AlibabawdksyncedorderqueryAPIRequest对象
func NewAlibabawdksyncedorderqueryRequest() *AlibabawdksyncedorderqueryAPIRequest {
	return &AlibabawdksyncedorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdksyncedorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.syncedorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdksyncedorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdksyncedorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabawdksyncedorderqueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdksyncedorderqueryAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSerialNum is SerialNum Setter
// 序列号
func (r *AlibabawdksyncedorderqueryAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// GetSerialNum SerialNum Getter
func (r AlibabawdksyncedorderqueryAPIRequest) GetSerialNum() string {
	return r._serialNum
}
