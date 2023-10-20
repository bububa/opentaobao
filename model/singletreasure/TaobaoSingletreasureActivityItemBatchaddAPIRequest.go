package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemBatchaddAPIRequest 批量添加商品接口 API请求
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
type TaobaoSingletreasureActivityItemBatchaddAPIRequest struct {
	model.Params
	// 系统入参
	_itemDetailInfo *ItemDetailInfoBatchCreateDto
}

// NewTaobaoSingletreasureActivityItemBatchaddRequest 初始化TaobaoSingletreasureActivityItemBatchaddAPIRequest对象
func NewTaobaoSingletreasureActivityItemBatchaddRequest() *TaobaoSingletreasureActivityItemBatchaddAPIRequest {
	return &TaobaoSingletreasureActivityItemBatchaddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityItemBatchaddAPIRequest) Reset() {
	r._itemDetailInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.batchadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDetailInfo is ItemDetailInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityItemBatchaddAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
	r._itemDetailInfo = _itemDetailInfo
	r.Set("item_detail_info", _itemDetailInfo)
	return nil
}

// GetItemDetailInfo ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
	return r._itemDetailInfo
}

var poolTaobaoSingletreasureActivityItemBatchaddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityItemBatchaddRequest()
	},
}

// GetTaobaoSingletreasureActivityItemBatchaddRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityItemBatchaddAPIRequest
func GetTaobaoSingletreasureActivityItemBatchaddAPIRequest() *TaobaoSingletreasureActivityItemBatchaddAPIRequest {
	return poolTaobaoSingletreasureActivityItemBatchaddAPIRequest.Get().(*TaobaoSingletreasureActivityItemBatchaddAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityItemBatchaddAPIRequest 将 TaobaoSingletreasureActivityItemBatchaddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemBatchaddAPIRequest(v *TaobaoSingletreasureActivityItemBatchaddAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemBatchaddAPIRequest.Put(v)
}
