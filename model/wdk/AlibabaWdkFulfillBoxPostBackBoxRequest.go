package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
RT收箱回传 APIRequest
alibaba.wdk.fulfill.box.post.back.box

RT收箱后，信息同步履约，履约同通知UMS 容器管理
*/
type AlibabaWdkFulfillBoxPostBackBoxRequest struct {
    model.Params

    // RT收箱回传请求参数
    returnBoxContainerRequest   *ReturnBoxContainerRequest 

}

func NewAlibabaWdkFulfillBoxPostBackBoxRequest() *AlibabaWdkFulfillBoxPostBackBoxRequest{
    return &AlibabaWdkFulfillBoxPostBackBoxRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillBoxPostBackBoxRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.box.post.back.box"
}

func (r AlibabaWdkFulfillBoxPostBackBoxRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillBoxPostBackBoxRequest) SetReturnBoxContainerRequest(returnBoxContainerRequest *ReturnBoxContainerRequest) error {
    r.returnBoxContainerRequest = returnBoxContainerRequest
    r.Set("return_box_container_request", returnBoxContainerRequest)
    return nil
}

func (r AlibabaWdkFulfillBoxPostBackBoxRequest) GetReturnBoxContainerRequest() *ReturnBoxContainerRequest {
    return r.returnBoxContainerRequest
}

