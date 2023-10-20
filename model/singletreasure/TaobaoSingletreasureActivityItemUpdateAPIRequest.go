package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitemupdateAPIRequest 更新单品优惠接口 API请求
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
type TaobaosingletreasureactivityitemupdateAPIRequest struct {
	model.Params
	// 修改接口的入参对象
	_itemDetailInfo *ItemDetailInfoCreateDto
}

// NewTaobaosingletreasureactivityitemupdateRequest 初始化TaobaosingletreasureactivityitemupdateAPIRequest对象
func NewTaobaosingletreasureactivityitemupdateRequest() *TaobaosingletreasureactivityitemupdateAPIRequest {
	return &TaobaosingletreasureactivityitemupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivityitemupdateAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivityitemupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivityitemupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDetailInfo is ItemDetailInfo Setter
// 修改接口的入参对象
func (r *TaobaosingletreasureactivityitemupdateAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoCreateDto) error {
	r._itemDetailInfo = _itemDetailInfo
	r.Set("item_detail_info", _itemDetailInfo)
	return nil
}

// GetItemDetailInfo ItemDetailInfo Getter
func (r TaobaosingletreasureactivityitemupdateAPIRequest) GetItemDetailInfo() *ItemDetailInfoCreateDto {
	return r._itemDetailInfo
}
