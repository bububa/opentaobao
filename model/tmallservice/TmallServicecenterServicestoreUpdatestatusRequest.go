package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网点/门店状态修改 API请求
tmall.servicecenter.servicestore.updatestatus

修改网点/门店状态
*/
type TmallServicecenterServicestoreUpdatestatusRequest struct {
    model.Params
    // 门店id
    _id   int64
    // 状态。1 营业，0歇业，-1彻底关店
    _status   int64
    // 业务类型。不同业务传不同的值
    _bizType   string
}

// 初始化TmallServicecenterServicestoreUpdatestatusRequest对象
func NewTmallServicecenterServicestoreUpdatestatusRequest() *TmallServicecenterServicestoreUpdatestatusRequest{
    return &TmallServicecenterServicestoreUpdatestatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.updatestatus"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 门店id
func (r *TmallServicecenterServicestoreUpdatestatusRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetId() int64 {
    return r._id
}
// Status Setter
// 状态。1 营业，0歇业，-1彻底关店
func (r *TmallServicecenterServicestoreUpdatestatusRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetStatus() int64 {
    return r._status
}
// BizType Setter
// 业务类型。不同业务传不同的值
func (r *TmallServicecenterServicestoreUpdatestatusRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetBizType() string {
    return r._bizType
}
