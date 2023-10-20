package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomtypeconflictdataAPIRequest 商家床型冲突数据接口 API请求
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
type TaobaoxhotelroomtypeconflictdataAPIRequest struct {
	model.Params
	// 查询第几页数据
	_currentPage int64
}

// NewTaobaoxhotelroomtypeconflictdataRequest 初始化TaobaoxhotelroomtypeconflictdataAPIRequest对象
func NewTaobaoxhotelroomtypeconflictdataRequest() *TaobaoxhotelroomtypeconflictdataAPIRequest {
	return &TaobaoxhotelroomtypeconflictdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelroomtypeconflictdataAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.conflict.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelroomtypeconflictdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelroomtypeconflictdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCurrentPage is CurrentPage Setter
// 查询第几页数据
func (r *TaobaoxhotelroomtypeconflictdataAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoxhotelroomtypeconflictdataAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
