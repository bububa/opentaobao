package singletreasure

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemBatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemBatchupdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
