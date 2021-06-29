package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合同下载 API请求
tmall.car.contract.download

目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限；
*/
type TmallCarContractDownloadRequest struct {
    model.Params
    // 天猫订单号
    orderId   int64
    // 是否下载html，true是html，false是pdf， html速度会快一点
    html   bool
}

// 初始化TmallCarContractDownloadRequest对象
func NewTmallCarContractDownloadRequest() *TmallCarContractDownloadRequest{
    return &TmallCarContractDownloadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarContractDownloadRequest) GetApiMethodName() string {
    return "tmall.car.contract.download"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarContractDownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫订单号
func (r *TmallCarContractDownloadRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarContractDownloadRequest) GetOrderId() int64 {
    return r.orderId
}
// Html Setter
// 是否下载html，true是html，false是pdf， html速度会快一点
func (r *TmallCarContractDownloadRequest) SetHtml(html bool) error {
    r.html = html
    r.Set("html", html)
    return nil
}

// Html Getter
func (r TmallCarContractDownloadRequest) GetHtml() bool {
    return r.html
}
