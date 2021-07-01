package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询制卡商制卡进度 API请求
alibaba.fundplatform.cardorder.status.query

当通知制卡商进行制卡后，其制卡流程会比较长，若长时间未反馈当前制卡进展，则需要使用该接口来向制卡商发起进度查询。
*/
type AlibabaFundplatformCardorderStatusQueryAPIRequest struct {
    model.Params
    // 请求结构体
    _request   *AlibabaFundplatformCardorderStatusQueryStruct
}

// 初始化AlibabaFundplatformCardorderStatusQueryAPIRequest对象
func NewAlibabaFundplatformCardorderStatusQueryRequest() *AlibabaFundplatformCardorderStatusQueryAPIRequest{
    return &AlibabaFundplatformCardorderStatusQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderStatusQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderStatusQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 请求结构体
func (r *AlibabaFundplatformCardorderStatusQueryAPIRequest) SetRequest(_request *AlibabaFundplatformCardorderStatusQueryStruct) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r AlibabaFundplatformCardorderStatusQueryAPIRequest) GetRequest() *AlibabaFundplatformCardorderStatusQueryStruct {
    return r._request
}
