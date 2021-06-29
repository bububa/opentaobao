package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元内绑定资源位 API请求
taobao.feedflow.item.adgroup.adzone.bind

信息流单元内绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzoneBindRequest struct {
    model.Params
    // 新增的绑定资源位
    _bindAdzoneList   []AdzoneBindDTO
    // 单元id
    _adgroupId   int64
}

// 初始化TaobaoFeedflowItemAdgroupAdzoneBindRequest对象
func NewTaobaoFeedflowItemAdgroupAdzoneBindRequest() *TaobaoFeedflowItemAdgroupAdzoneBindRequest{
    return &TaobaoFeedflowItemAdgroupAdzoneBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BindAdzoneList Setter
// 新增的绑定资源位
func (r *TaobaoFeedflowItemAdgroupAdzoneBindRequest) SetBindAdzoneList(_bindAdzoneList []AdzoneBindDTO) error {
    r._bindAdzoneList = _bindAdzoneList
    r.Set("bind_adzone_list", _bindAdzoneList)
    return nil
}

// BindAdzoneList Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetBindAdzoneList() []AdzoneBindDTO {
    return r._bindAdzoneList
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupAdzoneBindRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
