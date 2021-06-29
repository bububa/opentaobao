package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
款式生产信息同步API APIRequest
tmall.trend.style.produceinfo.upload

款式生产信息同步至平台
*/
type TmallTrendStyleProduceinfoUploadRequest struct {
    model.Params

    // 款式生产信息列表，单次同步最对1000条
    styleProduceInfoBoList   []StyleProduceInfoBO 

}

func NewTmallTrendStyleProduceinfoUploadRequest() *TmallTrendStyleProduceinfoUploadRequest{
    return &TmallTrendStyleProduceinfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTrendStyleProduceinfoUploadRequest) GetApiMethodName() string {
    return "tmall.trend.style.produceinfo.upload"
}

func (r TmallTrendStyleProduceinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTrendStyleProduceinfoUploadRequest) SetStyleProduceInfoBoList(styleProduceInfoBoList []StyleProduceInfoBO) error {
    r.styleProduceInfoBoList = styleProduceInfoBoList
    r.Set("style_produce_info_bo_list", styleProduceInfoBoList)
    return nil
}

func (r TmallTrendStyleProduceinfoUploadRequest) GetStyleProduceInfoBoList() []StyleProduceInfoBO {
    return r.styleProduceInfoBoList
}

