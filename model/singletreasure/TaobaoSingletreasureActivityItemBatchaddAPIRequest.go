package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitembatchaddAPIRequest 批量添加商品接口 API请求
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
type TaobaosingletreasureactivityitembatchaddAPIRequest struct {
	model.Params
	// 系统入参
	_itemDetailInfo *ItemDetailInfoBatchCreateDto
}

// NewTaobaosingletreasureactivityitembatchaddRequest 初始化TaobaosingletreasureactivityitembatchaddAPIRequest对象
func NewTaobaosingletreasureactivityitembatchaddRequest() *TaobaosingletreasureactivityitembatchaddAPIRequest {
	return &TaobaosingletreasureactivityitembatchaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivityitembatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.batchadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivityitembatchaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivityitembatchaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDetailInfo is ItemDetailInfo Setter
// 系统入参
func (r *TaobaosingletreasureactivityitembatchaddAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
	r._itemDetailInfo = _itemDetailInfo
	r.Set("item_detail_info", _itemDetailInfo)
	return nil
}

// GetItemDetailInfo ItemDetailInfo Getter
func (r TaobaosingletreasureactivityitembatchaddAPIRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
	return r._itemDetailInfo
}
