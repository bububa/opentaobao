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
type AlibabaXiamiApiSearchHotwordsGetAPIRequest struct {
    model.Params
    // 数量
    _limit   int64
}

// 初始化AlibabaXiamiApiSearchHotwordsGetAPIRequest对象
func NewAlibabaXiamiApiSearchHotwordsGetRequest() *AlibabaXiamiApiSearchHotwordsGetAPIRequest{
    return &AlibabaXiamiApiSearchHotwordsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiSearchHotwordsGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.search.hotwords.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiSearchHotwordsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Limit Setter
// 数量
func (r *AlibabaXiamiApiSearchHotwordsGetAPIRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiSearchHotwordsGetAPIRequest) GetLimit() int64 {
    return r._limit
}
