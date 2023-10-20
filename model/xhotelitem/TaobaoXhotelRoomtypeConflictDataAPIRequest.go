package xhotelitem

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRoomtypeConflictDataAPIRequest) Reset() {
	r._currentPage = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.conflict.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrentPage is CurrentPage Setter
// 查询第几页数据
func (r *TaobaoXhotelRoomtypeConflictDataAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoXhotelRoomtypeConflictDataAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

var poolTaobaoXhotelRoomtypeConflictDataAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRoomtypeConflictDataRequest()
	},
}

// GetTaobaoXhotelRoomtypeConflictDataRequest 从 sync.Pool 获取 TaobaoXhotelRoomtypeConflictDataAPIRequest
func GetTaobaoXhotelRoomtypeConflictDataAPIRequest() *TaobaoXhotelRoomtypeConflictDataAPIRequest {
	return poolTaobaoXhotelRoomtypeConflictDataAPIRequest.Get().(*TaobaoXhotelRoomtypeConflictDataAPIRequest)
}

// ReleaseTaobaoXhotelRoomtypeConflictDataAPIRequest 将 TaobaoXhotelRoomtypeConflictDataAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRoomtypeConflictDataAPIRequest(v *TaobaoXhotelRoomtypeConflictDataAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRoomtypeConflictDataAPIRequest.Put(v)
}
