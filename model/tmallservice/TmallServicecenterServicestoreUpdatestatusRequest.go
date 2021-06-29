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
    id   int64
    // 状态。1 营业，0歇业，-1彻底关店
    status   int64
    // 业务类型。不同业务传不同的值
    bizType   string
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
func (r *TmallServicecenterServicestoreUpdatestatusRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetId() int64 {
    return r.id
}
// Status Setter
// 状态。1 营业，0歇业，-1彻底关店
func (r *TmallServicecenterServicestoreUpdatestatusRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetStatus() int64 {
    return r.status
}
// BizType Setter
// 业务类型。不同业务传不同的值
func (r *TmallServicecenterServicestoreUpdatestatusRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterServicestoreUpdatestatusRequest) GetBizType() string {
    return r.bizType
}
