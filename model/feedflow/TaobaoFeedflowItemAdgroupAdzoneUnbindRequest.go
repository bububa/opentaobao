package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元内解绑资源位 API请求
taobao.feedflow.item.adgroup.adzone.unbind

信息流单元内解绑资源位
*/
type TaobaoFeedflowItemAdgroupAdzoneUnbindRequest struct {
    model.Params
    // 广告位id
    _adzoneIdList   []int64
    // 单元id
    _adgroupId   int64
}

// 初始化TaobaoFeedflowItemAdgroupAdzoneUnbindRequest对象
func NewTaobaoFeedflowItemAdgroupAdzoneUnbindRequest() *TaobaoFeedflowItemAdgroupAdzoneUnbindRequest{
    return &TaobaoFeedflowItemAdgroupAdzoneUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.unbind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdzoneIdList Setter
// 广告位id
func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) SetAdzoneIdList(_adzoneIdList []int64) error {
    r._adzoneIdList = _adzoneIdList
    r.Set("adzone_id_list", _adzoneIdList)
    return nil
}

// AdzoneIdList Getter
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetAdzoneIdList() []int64 {
    return r._adzoneIdList
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
