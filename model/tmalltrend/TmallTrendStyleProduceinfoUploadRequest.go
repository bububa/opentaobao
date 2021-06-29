package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
款式生产信息同步API API请求
tmall.trend.style.produceinfo.upload

款式生产信息同步至平台
*/
type TmallTrendStyleProduceinfoUploadRequest struct {
    model.Params
    // 款式生产信息列表，单次同步最对1000条
    _styleProduceInfoBoList   []StyleProduceInfoBO
}

// 初始化TmallTrendStyleProduceinfoUploadRequest对象
func NewTmallTrendStyleProduceinfoUploadRequest() *TmallTrendStyleProduceinfoUploadRequest{
    return &TmallTrendStyleProduceinfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTrendStyleProduceinfoUploadRequest) GetApiMethodName() string {
    return "tmall.trend.style.produceinfo.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTrendStyleProduceinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StyleProduceInfoBoList Setter
// 款式生产信息列表，单次同步最对1000条
func (r *TmallTrendStyleProduceinfoUploadRequest) SetStyleProduceInfoBoList(_styleProduceInfoBoList []StyleProduceInfoBO) error {
    r._styleProduceInfoBoList = _styleProduceInfoBoList
    r.Set("style_produce_info_bo_list", _styleProduceInfoBoList)
    return nil
}

// StyleProduceInfoBoList Getter
func (r TmallTrendStyleProduceinfoUploadRequest) GetStyleProduceInfoBoList() []StyleProduceInfoBO {
    return r._styleProduceInfoBoList
}
