package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元下查看绑定资源位 API请求
taobao.feedflow.item.adgroup.adzone.page

信息流单元下查看绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzonePageAPIRequest struct {
    model.Params
    // 查询条件
    _adzoneBindQuery   *AdzoneBindQueryDTO
}

// 初始化TaobaoFeedflowItemAdgroupAdzonePageAPIRequest对象
func NewTaobaoFeedflowItemAdgroupAdzonePageRequest() *TaobaoFeedflowItemAdgroupAdzonePageAPIRequest{
    return &TaobaoFeedflowItemAdgroupAdzonePageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdzoneBindQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) SetAdzoneBindQuery(_adzoneBindQuery *AdzoneBindQueryDTO) error {
    r._adzoneBindQuery = _adzoneBindQuery
    r.Set("adzone_bind_query", _adzoneBindQuery)
    return nil
}

// AdzoneBindQuery Getter
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetAdzoneBindQuery() *AdzoneBindQueryDTO {
    return r._adzoneBindQuery
}
