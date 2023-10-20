package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkopencateorderpullAPIRequest 商户回传餐饮加工单状态 API请求
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
type AlibabawdkopencateorderpullAPIRequest struct {
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

// NewAlibabawdkopencateorderpullRequest 初始化AlibabawdkopencateorderpullAPIRequest对象
func NewAlibabawdkopencateorderpullRequest() *AlibabawdkopencateorderpullAPIRequest {
	return &AlibabawdkopencateorderpullAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkopencateorderpullAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkopen.cateorder.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkopencateorderpullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkopencateorderpullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOutOrderIds is SubOutOrderIds Setter
// 主站子订单ID列表, 为空则表示回传整单状态
func (r *AlibabawdkopencateorderpullAPIRequest) SetSubOutOrderIds(_subOutOrderIds []string) error {
	r._subOutOrderIds = _subOutOrderIds
	r.Set("sub_out_order_ids", _subOutOrderIds)
	return nil
}

// GetSubOutOrderIds SubOutOrderIds Getter
func (r AlibabawdkopencateorderpullAPIRequest) GetSubOutOrderIds() []string {
	return r._subOutOrderIds
}

// SetStoreId is StoreId Setter
// 经营店ID
func (r *AlibabawdkopencateorderpullAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkopencateorderpullAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 主站主订单ID
func (r *AlibabawdkopencateorderpullAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabawdkopencateorderpullAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetStatus is Status Setter
// 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
func (r *AlibabawdkopencateorderpullAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabawdkopencateorderpullAPIRequest) GetStatus() string {
	return r._status
}
