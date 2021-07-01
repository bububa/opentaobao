package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售商家仓接单 API请求
alibaba.ascp.uop.taobao.presalesorder.create

预售商家仓接单
*/
type AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest struct {
    model.Params
    // 预售商家仓接单对象
    _presalesOrderCreateRequest   *PresalesordercreaterequestTest
}

// 初始化AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest对象
func NewAlibabaAscpUopTaobaoPresalesorderCreateRequest() *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest{
    return &AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.taobao.presalesorder.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PresalesOrderCreateRequest Setter
// 预售商家仓接单对象
func (r *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) SetPresalesOrderCreateRequest(_presalesOrderCreateRequest *PresalesordercreaterequestTest) error {
    r._presalesOrderCreateRequest = _presalesOrderCreateRequest
    r.Set("presales_order_create_request", _presalesOrderCreateRequest)
    return nil
}

// PresalesOrderCreateRequest Getter
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetPresalesOrderCreateRequest() *PresalesordercreaterequestTest {
    return r._presalesOrderCreateRequest
}
