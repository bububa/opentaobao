package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
穷游内容写入接口 APIRequest
alitrip.platform.content.raw.add

穷游内容写入飞猪接口
*/
type AlitripPlatformContentRawAddRequest struct {
    model.Params

    // 写入入参
    fliggyContentRequest   *FliggyContentRequest 

}

func NewAlitripPlatformContentRawAddRequest() *AlitripPlatformContentRawAddRequest{
    return &AlitripPlatformContentRawAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformContentRawAddRequest) GetApiMethodName() string {
    return "alitrip.platform.content.raw.add"
}

func (r AlitripPlatformContentRawAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformContentRawAddRequest) SetFliggyContentRequest(fliggyContentRequest *FliggyContentRequest) error {
    r.fliggyContentRequest = fliggyContentRequest
    r.Set("fliggy_content_request", fliggyContentRequest)
    return nil
}

func (r AlitripPlatformContentRawAddRequest) GetFliggyContentRequest() *FliggyContentRequest {
    return r.fliggyContentRequest
}

