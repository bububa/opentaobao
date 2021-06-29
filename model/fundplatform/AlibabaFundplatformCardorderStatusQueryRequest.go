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
type AlibabaFundplatformCardorderStatusQueryRequest struct {
    model.Params
    // 请求结构体
    request   *AlibabaFundplatformCardorderStatusQueryStruct
}

// 初始化AlibabaFundplatformCardorderStatusQueryRequest对象
func NewAlibabaFundplatformCardorderStatusQueryRequest() *AlibabaFundplatformCardorderStatusQueryRequest{
    return &AlibabaFundplatformCardorderStatusQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 请求结构体
func (r *AlibabaFundplatformCardorderStatusQueryRequest) SetRequest(request *AlibabaFundplatformCardorderStatusQueryStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r AlibabaFundplatformCardorderStatusQueryRequest) GetRequest() *AlibabaFundplatformCardorderStatusQueryStruct {
    return r.request
}
