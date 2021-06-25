package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索热词 APIRequest
alibaba.xiami.api.search.hotwords.get

搜索热词
*/
type AlibabaXiamiApiSearchHotwordsGetRequest struct {
    model.Params

    // 数量
    limit   int64 

}

func NewAlibabaXiamiApiSearchHotwordsGetRequest() *AlibabaXiamiApiSearchHotwordsGetRequest{
    return &AlibabaXiamiApiSearchHotwordsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiSearchHotwordsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.search.hotwords.get"
}

func (r AlibabaXiamiApiSearchHotwordsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiSearchHotwordsGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiSearchHotwordsGetRequest) GetLimit() int64 {
    return r.limit
}

