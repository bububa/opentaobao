package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服评价详情接口 API请求
taobao.qianniu.kefueval.get

获取买家对客服的服务评价
*/
type TaobaoQianniuKefuevalGetAPIRequest struct {
    model.Params
    // 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
    _queryIds   string
    // 开始时间，格式yyyyMMddHHmmss
    _btime   string
    // 结束时间，格式yyyyMMddHHmmss
    _etime   string
}

// 初始化TaobaoQianniuKefuevalGetAPIRequest对象
func NewTaobaoQianniuKefuevalGetRequest() *TaobaoQianniuKefuevalGetAPIRequest{
    return &TaobaoQianniuKefuevalGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuKefuevalGetAPIRequest) GetApiMethodName() string {
    return "taobao.qianniu.kefueval.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuKefuevalGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryIds Setter
// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
func (r *TaobaoQianniuKefuevalGetAPIRequest) SetQueryIds(_queryIds string) error {
    r._queryIds = _queryIds
    r.Set("query_ids", _queryIds)
    return nil
}

// QueryIds Getter
func (r TaobaoQianniuKefuevalGetAPIRequest) GetQueryIds() string {
    return r._queryIds
}
// Btime Setter
// 开始时间，格式yyyyMMddHHmmss
func (r *TaobaoQianniuKefuevalGetAPIRequest) SetBtime(_btime string) error {
    r._btime = _btime
    r.Set("btime", _btime)
    return nil
}

// Btime Getter
func (r TaobaoQianniuKefuevalGetAPIRequest) GetBtime() string {
    return r._btime
}
// Etime Setter
// 结束时间，格式yyyyMMddHHmmss
func (r *TaobaoQianniuKefuevalGetAPIRequest) SetEtime(_etime string) error {
    r._etime = _etime
    r.Set("etime", _etime)
    return nil
}

// Etime Getter
func (r TaobaoQianniuKefuevalGetAPIRequest) GetEtime() string {
    return r._etime
}
