package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流新增并且绑定创意 API请求
taobao.feedflow.item.adgroup.creative.add.bind

信息流新增并且绑定创意
*/
type TaobaoFeedflowItemAdgroupCreativeAddBindRequest struct {
    model.Params
    // 新增绑定的创意，一次最多2个
    creativeBindList   []CreativeBindDto
    // 单元id
    adgroupId   int64
}

// 初始化TaobaoFeedflowItemAdgroupCreativeAddBindRequest对象
func NewTaobaoFeedflowItemAdgroupCreativeAddBindRequest() *TaobaoFeedflowItemAdgroupCreativeAddBindRequest{
    return &TaobaoFeedflowItemAdgroupCreativeAddBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.creative.add.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreativeBindList Setter
// 新增绑定的创意，一次最多2个
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindRequest) SetCreativeBindList(creativeBindList []CreativeBindDto) error {
    r.creativeBindList = creativeBindList
    r.Set("creative_bind_list", creativeBindList)
    return nil
}

// CreativeBindList Getter
func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetCreativeBindList() []CreativeBindDto {
    return r.creativeBindList
}
// AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupCreativeAddBindRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
