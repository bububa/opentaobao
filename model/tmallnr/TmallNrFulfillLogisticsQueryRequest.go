package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定时送和极速达配送物流信息查询 API请求
tmall.nr.fulfill.logistics.query

发布定时送&极速达物流信息查询服务
*/
type TmallNrFulfillLogisticsQueryRequest struct {
    model.Params
    // 交易主订单号
    mainOrderId   int64
    // 业务标识，dss标识定时送业务；jsd表示极速达业务
    bizIdentity   string
}

// 初始化TmallNrFulfillLogisticsQueryRequest对象
func NewTmallNrFulfillLogisticsQueryRequest() *TmallNrFulfillLogisticsQueryRequest{
    return &TmallNrFulfillLogisticsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillLogisticsQueryRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.logistics.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillLogisticsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 交易主订单号
func (r *TmallNrFulfillLogisticsQueryRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TmallNrFulfillLogisticsQueryRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrFulfillLogisticsQueryRequest) SetBizIdentity(bizIdentity string) error {
    r.bizIdentity = bizIdentity
    r.Set("biz_identity", bizIdentity)
    return nil
}

// BizIdentity Getter
func (r TmallNrFulfillLogisticsQueryRequest) GetBizIdentity() string {
    return r.bizIdentity
}
