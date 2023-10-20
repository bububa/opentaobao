package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemBatchupdateAPIRequest 批量修改商品接口 API请求
// taobao.singletreasure.activity.item.batchupdate
//
// 批量修改商品优惠接口
type TaobaoSingletreasureActivityItemBatchupdateAPIRequest struct {
	model.Params
	// 系统入参
	_itemDetailInfo *ItemDetailInfoBatchCreateDto
}

// NewTaobaoSingletreasureActivityItemBatchupdateRequest 初始化TaobaoSingletreasureActivityItemBatchupdateAPIRequest对象
func NewTaobaoSingletreasureActivityItemBatchupdateRequest() *TaobaoSingletreasureActivityItemBatchupdateAPIRequest {
	return &TaobaoSingletreasureActivityItemBatchupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityItemBatchupdateAPIRequest) Reset() {
	r._itemDetailInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemBatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemBatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityItemBatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDetailInfo is ItemDetailInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityItemBatchupdateAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
	r._itemDetailInfo = _itemDetailInfo
	r.Set("item_detail_info", _itemDetailInfo)
	return nil
}

// GetItemDetailInfo ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemBatchupdateAPIRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
	return r._itemDetailInfo
}

var poolTaobaoSingletreasureActivityItemBatchupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityItemBatchupdateRequest()
	},
}

// GetTaobaoSingletreasureActivityItemBatchupdateRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityItemBatchupdateAPIRequest
func GetTaobaoSingletreasureActivityItemBatchupdateAPIRequest() *TaobaoSingletreasureActivityItemBatchupdateAPIRequest {
	return poolTaobaoSingletreasureActivityItemBatchupdateAPIRequest.Get().(*TaobaoSingletreasureActivityItemBatchupdateAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityItemBatchupdateAPIRequest 将 TaobaoSingletreasureActivityItemBatchupdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemBatchupdateAPIRequest(v *TaobaoSingletreasureActivityItemBatchupdateAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemBatchupdateAPIRequest.Put(v)
}
