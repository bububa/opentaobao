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
    bindAdzoneList   []AdzoneBindDto
    // 单元id
    adgroupId   int64
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
func (r *TaobaoFeedflowItemAdgroupAdzoneBindRequest) SetBindAdzoneList(bindAdzoneList []AdzoneBindDto) error {
    r.bindAdzoneList = bindAdzoneList
    r.Set("bind_adzone_list", bindAdzoneList)
    return nil
}

// BindAdzoneList Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetBindAdzoneList() []AdzoneBindDto {
    return r.bindAdzoneList
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupAdzoneBindRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
