package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家床型冲突数据接口 API请求
taobao.xhotel.roomtype.conflict.data

商家床型冲突数据接口
*/
type TaobaoXhotelRoomtypeConflictDataRequest struct {
    model.Params
    // 查询第几页数据
    currentPage   int64
}

// 初始化TaobaoXhotelRoomtypeConflictDataRequest对象
func NewTaobaoXhotelRoomtypeConflictDataRequest() *TaobaoXhotelRoomtypeConflictDataRequest{
    return &TaobaoXhotelRoomtypeConflictDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeConflictDataRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.conflict.data"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeConflictDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// 查询第几页数据
func (r *TaobaoXhotelRoomtypeConflictDataRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoXhotelRoomtypeConflictDataRequest) GetCurrentPage() int64 {
    return r.currentPage
}
