package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户是否收藏某个店铺 API请求
alibaba.interact.open.isattention

判断用户是否收藏某个店铺
*/
type AlibabaInteractOpenIsattentionRequest struct {
    model.Params
    // 1
    param0   int64
}

// 初始化AlibabaInteractOpenIsattentionRequest对象
func NewAlibabaInteractOpenIsattentionRequest() *AlibabaInteractOpenIsattentionRequest{
    return &AlibabaInteractOpenIsattentionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractOpenIsattentionRequest) GetApiMethodName() string {
    return "alibaba.interact.open.isattention"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractOpenIsattentionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 1
func (r *AlibabaInteractOpenIsattentionRequest) SetParam0(param0 int64) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaInteractOpenIsattentionRequest) GetParam0() int64 {
    return r.param0
}
