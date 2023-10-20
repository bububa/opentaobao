package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemUpdateAPIRequest 更新单品优惠接口 API请求
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
type TaobaoSingletreasureActivityItemUpdateAPIRequest struct {
	model.Params
	// 修改接口的入参对象
	_itemDetailInfo *ItemDetailInfoCreateDto
}

// NewTaobaoSingletreasureActivityItemUpdateRequest 初始化TaobaoSingletreasureActivityItemUpdateAPIRequest对象
func NewTaobaoSingletreasureActivityItemUpdateRequest() *TaobaoSingletreasureActivityItemUpdateAPIRequest {
	return &TaobaoSingletreasureActivityItemUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityItemUpdateAPIRequest) Reset() {
	r._itemDetailInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDetailInfo is ItemDetailInfo Setter
// 修改接口的入参对象
func (r *TaobaoSingletreasureActivityItemUpdateAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoCreateDto) error {
	r._itemDetailInfo = _itemDetailInfo
	r.Set("item_detail_info", _itemDetailInfo)
	return nil
}

// GetItemDetailInfo ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetItemDetailInfo() *ItemDetailInfoCreateDto {
	return r._itemDetailInfo
}

var poolTaobaoSingletreasureActivityItemUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityItemUpdateRequest()
	},
}

// GetTaobaoSingletreasureActivityItemUpdateRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityItemUpdateAPIRequest
func GetTaobaoSingletreasureActivityItemUpdateAPIRequest() *TaobaoSingletreasureActivityItemUpdateAPIRequest {
	return poolTaobaoSingletreasureActivityItemUpdateAPIRequest.Get().(*TaobaoSingletreasureActivityItemUpdateAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityItemUpdateAPIRequest 将 TaobaoSingletreasureActivityItemUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemUpdateAPIRequest(v *TaobaoSingletreasureActivityItemUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemUpdateAPIRequest.Put(v)
}
