package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallGoodsStatusSyncAPIRequest 第三方商城接入采购商城-商品状态同步 API请求
// alibaba.pur.cmall.goods.status.sync
//
// 第三方商城接入采购商城-商品状态同步
type AlibabaPurCmallGoodsStatusSyncAPIRequest struct {
	model.Params
	// 外部商家标识
	_goodsDataSource string
	// 外部商家商品唯一标志
	_goodsSourceValue string
	// 状态，shelfed:上架，unShelfed:下架
	_status string
}

// NewAlibabaPurCmallGoodsStatusSyncRequest 初始化AlibabaPurCmallGoodsStatusSyncAPIRequest对象
func NewAlibabaPurCmallGoodsStatusSyncRequest() *AlibabaPurCmallGoodsStatusSyncAPIRequest {
	return &AlibabaPurCmallGoodsStatusSyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurCmallGoodsStatusSyncAPIRequest) Reset() {
	r._goodsDataSource = ""
	r._goodsSourceValue = ""
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurCmallGoodsStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.cmall.goods.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurCmallGoodsStatusSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurCmallGoodsStatusSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDataSource is GoodsDataSource Setter
// 外部商家标识
func (r *AlibabaPurCmallGoodsStatusSyncAPIRequest) SetGoodsDataSource(_goodsDataSource string) error {
	r._goodsDataSource = _goodsDataSource
	r.Set("goods_data_source", _goodsDataSource)
	return nil
}

// GetGoodsDataSource GoodsDataSource Getter
func (r AlibabaPurCmallGoodsStatusSyncAPIRequest) GetGoodsDataSource() string {
	return r._goodsDataSource
}

// SetGoodsSourceValue is GoodsSourceValue Setter
// 外部商家商品唯一标志
func (r *AlibabaPurCmallGoodsStatusSyncAPIRequest) SetGoodsSourceValue(_goodsSourceValue string) error {
	r._goodsSourceValue = _goodsSourceValue
	r.Set("goods_source_value", _goodsSourceValue)
	return nil
}

// GetGoodsSourceValue GoodsSourceValue Getter
func (r AlibabaPurCmallGoodsStatusSyncAPIRequest) GetGoodsSourceValue() string {
	return r._goodsSourceValue
}

// SetStatus is Status Setter
// 状态，shelfed:上架，unShelfed:下架
func (r *AlibabaPurCmallGoodsStatusSyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaPurCmallGoodsStatusSyncAPIRequest) GetStatus() string {
	return r._status
}

var poolAlibabaPurCmallGoodsStatusSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurCmallGoodsStatusSyncRequest()
	},
}

// GetAlibabaPurCmallGoodsStatusSyncRequest 从 sync.Pool 获取 AlibabaPurCmallGoodsStatusSyncAPIRequest
func GetAlibabaPurCmallGoodsStatusSyncAPIRequest() *AlibabaPurCmallGoodsStatusSyncAPIRequest {
	return poolAlibabaPurCmallGoodsStatusSyncAPIRequest.Get().(*AlibabaPurCmallGoodsStatusSyncAPIRequest)
}

// ReleaseAlibabaPurCmallGoodsStatusSyncAPIRequest 将 AlibabaPurCmallGoodsStatusSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurCmallGoodsStatusSyncAPIRequest(v *AlibabaPurCmallGoodsStatusSyncAPIRequest) {
	v.Reset()
	poolAlibabaPurCmallGoodsStatusSyncAPIRequest.Put(v)
}
