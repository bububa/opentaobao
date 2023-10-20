package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkopenCateorderPullAPIRequest 商户回传餐饮加工单状态 API请求
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
type AlibabaWdkopenCateorderPullAPIRequest struct {
	model.Params
	// 主站子订单ID列表, 为空则表示回传整单状态
	_subOutOrderIds []string
	// 经营店ID
	_storeId string
	// 主站主订单ID
	_outOrderId string
	// 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
	_status string
}

// NewAlibabaWdkopenCateorderPullRequest 初始化AlibabaWdkopenCateorderPullAPIRequest对象
func NewAlibabaWdkopenCateorderPullRequest() *AlibabaWdkopenCateorderPullAPIRequest {
	return &AlibabaWdkopenCateorderPullAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkopenCateorderPullAPIRequest) Reset() {
	r._subOutOrderIds = r._subOutOrderIds[:0]
	r._storeId = ""
	r._outOrderId = ""
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkopenCateorderPullAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkopen.cateorder.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkopenCateorderPullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkopenCateorderPullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOutOrderIds is SubOutOrderIds Setter
// 主站子订单ID列表, 为空则表示回传整单状态
func (r *AlibabaWdkopenCateorderPullAPIRequest) SetSubOutOrderIds(_subOutOrderIds []string) error {
	r._subOutOrderIds = _subOutOrderIds
	r.Set("sub_out_order_ids", _subOutOrderIds)
	return nil
}

// GetSubOutOrderIds SubOutOrderIds Getter
func (r AlibabaWdkopenCateorderPullAPIRequest) GetSubOutOrderIds() []string {
	return r._subOutOrderIds
}

// SetStoreId is StoreId Setter
// 经营店ID
func (r *AlibabaWdkopenCateorderPullAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkopenCateorderPullAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 主站主订单ID
func (r *AlibabaWdkopenCateorderPullAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaWdkopenCateorderPullAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetStatus is Status Setter
// 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
func (r *AlibabaWdkopenCateorderPullAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaWdkopenCateorderPullAPIRequest) GetStatus() string {
	return r._status
}

var poolAlibabaWdkopenCateorderPullAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkopenCateorderPullRequest()
	},
}

// GetAlibabaWdkopenCateorderPullRequest 从 sync.Pool 获取 AlibabaWdkopenCateorderPullAPIRequest
func GetAlibabaWdkopenCateorderPullAPIRequest() *AlibabaWdkopenCateorderPullAPIRequest {
	return poolAlibabaWdkopenCateorderPullAPIRequest.Get().(*AlibabaWdkopenCateorderPullAPIRequest)
}

// ReleaseAlibabaWdkopenCateorderPullAPIRequest 将 AlibabaWdkopenCateorderPullAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkopenCateorderPullAPIRequest(v *AlibabaWdkopenCateorderPullAPIRequest) {
	v.Reset()
	poolAlibabaWdkopenCateorderPullAPIRequest.Put(v)
}
