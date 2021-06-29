package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方车型数据上传 API请求
tmall.aliauto.wisdomdata.model.recive

天猫汽车对外提供的汽车车型数据上传接口
*/
type TmallAliautoWisdomdataModelReciveRequest struct {
    model.Params
    // JSON格式车型完整数据
    modelDetail   string
    // 接入的第三方库中的车型唯一id
    resourceId   string
    // 接入的第三方库的名称
    fromSource   string
}

// 初始化TmallAliautoWisdomdataModelReciveRequest对象
func NewTmallAliautoWisdomdataModelReciveRequest() *TmallAliautoWisdomdataModelReciveRequest{
    return &TmallAliautoWisdomdataModelReciveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoWisdomdataModelReciveRequest) GetApiMethodName() string {
    return "tmall.aliauto.wisdomdata.model.recive"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoWisdomdataModelReciveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModelDetail Setter
// JSON格式车型完整数据
func (r *TmallAliautoWisdomdataModelReciveRequest) SetModelDetail(modelDetail string) error {
    r.modelDetail = modelDetail
    r.Set("model_detail", modelDetail)
    return nil
}

// ModelDetail Getter
func (r TmallAliautoWisdomdataModelReciveRequest) GetModelDetail() string {
    return r.modelDetail
}
// ResourceId Setter
// 接入的第三方库中的车型唯一id
func (r *TmallAliautoWisdomdataModelReciveRequest) SetResourceId(resourceId string) error {
    r.resourceId = resourceId
    r.Set("resource_id", resourceId)
    return nil
}

// ResourceId Getter
func (r TmallAliautoWisdomdataModelReciveRequest) GetResourceId() string {
    return r.resourceId
}
// FromSource Setter
// 接入的第三方库的名称
func (r *TmallAliautoWisdomdataModelReciveRequest) SetFromSource(fromSource string) error {
    r.fromSource = fromSource
    r.Set("from_source", fromSource)
    return nil
}

// FromSource Getter
func (r TmallAliautoWisdomdataModelReciveRequest) GetFromSource() string {
    return r.fromSource
}
