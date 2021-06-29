package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索热词 API请求
alibaba.xiami.api.search.hotwords.get

搜索热词
*/
type AlibabaXiamiApiSearchHotwordsGetRequest struct {
    model.Params
    // 数量
    _limit   int64
}

// 初始化AlibabaXiamiApiSearchHotwordsGetRequest对象
func NewAlibabaXiamiApiSearchHotwordsGetRequest() *AlibabaXiamiApiSearchHotwordsGetRequest{
    return &AlibabaXiamiApiSearchHotwordsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiSearchHotwordsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.search.hotwords.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiSearchHotwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Limit Setter
// 数量
func (r *AlibabaXiamiApiSearchHotwordsGetRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiSearchHotwordsGetRequest) GetLimit() int64 {
    return r._limit
}
