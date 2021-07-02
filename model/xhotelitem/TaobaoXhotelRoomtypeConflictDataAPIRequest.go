package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeConflictDataAPIRequest 商家床型冲突数据接口 API请求
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
type TaobaoXhotelRoomtypeConflictDataAPIRequest struct {
	model.Params
	// 查询第几页数据
	_currentPage int64
}

// NewTaobaoXhotelRoomtypeConflictDataRequest 初始化TaobaoXhotelRoomtypeConflictDataAPIRequest对象
func NewTaobaoXhotelRoomtypeConflictDataRequest() *TaobaoXhotelRoomtypeConflictDataAPIRequest {
	return &TaobaoXhotelRoomtypeConflictDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.conflict.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CurrentPage Setter
// 查询第几页数据
func (r *TaobaoXhotelRoomtypeConflictDataAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
