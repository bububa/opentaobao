package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityItemBatchaddAPIRequest
批量添加商品接口 API请求
taobao.singletreasure.activity.item.batchadd

向活动中批量添加商品优惠 */
type TaobaoSingletreasureActivityItemBatchaddAPIRequest struct {
	model.Params
	// 系统入参
	_itemDetailInfo *ItemDetailInfoBatchCreateDto
}

// NewTaobaoSingletreasureActivityItemBatchaddRequest 初始化TaobaoSingletreasureActivityItemBatchaddAPIRequest对象
func NewTaobaoSingletreasureActivityItemBatchaddRequest() *TaobaoSingletreasureActivityItemBatchaddAPIRequest {
	return &TaobaoSingletreasureActivityItemBatchaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.batchadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemDetailInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityItemBatchaddAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
	r._itemDetailInfo = _itemDetailInfo
	r.Set("item_detail_info", _itemDetailInfo)
	return nil
}

// Get ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemBatchaddAPIRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
	return r._itemDetailInfo
}
