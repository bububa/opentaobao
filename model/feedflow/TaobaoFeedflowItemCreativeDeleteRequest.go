package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流删除创意 API请求
taobao.feedflow.item.creative.delete

信息流删除创意
*/
type TaobaoFeedflowItemCreativeDeleteRequest struct {
    model.Params
    // 创意id列表
    _creativeIdList   []int64
}

// 初始化TaobaoFeedflowItemCreativeDeleteRequest对象
func NewTaobaoFeedflowItemCreativeDeleteRequest() *TaobaoFeedflowItemCreativeDeleteRequest{
    return &TaobaoFeedflowItemCreativeDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreativeIdList Setter
// 创意id列表
func (r *TaobaoFeedflowItemCreativeDeleteRequest) SetCreativeIdList(_creativeIdList []int64) error {
    r._creativeIdList = _creativeIdList
    r.Set("creative_id_list", _creativeIdList)
    return nil
}

// CreativeIdList Getter
func (r TaobaoFeedflowItemCreativeDeleteRequest) GetCreativeIdList() []int64 {
    return r._creativeIdList
}
