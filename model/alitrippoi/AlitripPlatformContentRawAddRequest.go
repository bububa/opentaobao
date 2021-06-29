package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
穷游内容写入接口 API请求
alitrip.platform.content.raw.add

穷游内容写入飞猪接口
*/
type AlitripPlatformContentRawAddRequest struct {
    model.Params
    // 写入入参
    _fliggyContentRequest   *FliggyContentRequest
}

// 初始化AlitripPlatformContentRawAddRequest对象
func NewAlitripPlatformContentRawAddRequest() *AlitripPlatformContentRawAddRequest{
    return &AlitripPlatformContentRawAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformContentRawAddRequest) GetApiMethodName() string {
    return "alitrip.platform.content.raw.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformContentRawAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FliggyContentRequest Setter
// 写入入参
func (r *AlitripPlatformContentRawAddRequest) SetFliggyContentRequest(_fliggyContentRequest *FliggyContentRequest) error {
    r._fliggyContentRequest = _fliggyContentRequest
    r.Set("fliggy_content_request", _fliggyContentRequest)
    return nil
}

// FliggyContentRequest Getter
func (r AlitripPlatformContentRawAddRequest) GetFliggyContentRequest() *FliggyContentRequest {
    return r._fliggyContentRequest
}
