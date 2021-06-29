package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方车型数据上传 APIRequest
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

func NewTmallAliautoWisdomdataModelReciveRequest() *TmallAliautoWisdomdataModelReciveRequest{
    return &TmallAliautoWisdomdataModelReciveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAliautoWisdomdataModelReciveRequest) GetApiMethodName() string {
    return "tmall.aliauto.wisdomdata.model.recive"
}

func (r TmallAliautoWisdomdataModelReciveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAliautoWisdomdataModelReciveRequest) SetModelDetail(modelDetail string) error {
    r.modelDetail = modelDetail
    r.Set("model_detail", modelDetail)
    return nil
}

func (r TmallAliautoWisdomdataModelReciveRequest) GetModelDetail() string {
    return r.modelDetail
}

func (r *TmallAliautoWisdomdataModelReciveRequest) SetResourceId(resourceId string) error {
    r.resourceId = resourceId
    r.Set("resource_id", resourceId)
    return nil
}

func (r TmallAliautoWisdomdataModelReciveRequest) GetResourceId() string {
    return r.resourceId
}

func (r *TmallAliautoWisdomdataModelReciveRequest) SetFromSource(fromSource string) error {
    r.fromSource = fromSource
    r.Set("from_source", fromSource)
    return nil
}

func (r TmallAliautoWisdomdataModelReciveRequest) GetFromSource() string {
    return r.fromSource
}

