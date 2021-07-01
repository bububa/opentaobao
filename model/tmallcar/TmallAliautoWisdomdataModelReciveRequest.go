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
type TmallAliautoWisdomdataModelReciveAPIRequest struct {
    model.Params
    // JSON格式车型完整数据
    _modelDetail   string
    // 接入的第三方库中的车型唯一id
    _resourceId   string
    // 接入的第三方库的名称
    _fromSource   string
}

// 初始化TmallAliautoWisdomdataModelReciveAPIRequest对象
func NewTmallAliautoWisdomdataModelReciveRequest() *TmallAliautoWisdomdataModelReciveAPIRequest{
    return &TmallAliautoWisdomdataModelReciveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetApiMethodName() string {
    return "tmall.aliauto.wisdomdata.model.recive"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModelDetail Setter
// JSON格式车型完整数据
func (r *TmallAliautoWisdomdataModelReciveAPIRequest) SetModelDetail(_modelDetail string) error {
    r._modelDetail = _modelDetail
    r.Set("model_detail", _modelDetail)
    return nil
}

// ModelDetail Getter
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetModelDetail() string {
    return r._modelDetail
}
// ResourceId Setter
// 接入的第三方库中的车型唯一id
func (r *TmallAliautoWisdomdataModelReciveAPIRequest) SetResourceId(_resourceId string) error {
    r._resourceId = _resourceId
    r.Set("resource_id", _resourceId)
    return nil
}

// ResourceId Getter
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetResourceId() string {
    return r._resourceId
}
// FromSource Setter
// 接入的第三方库的名称
func (r *TmallAliautoWisdomdataModelReciveAPIRequest) SetFromSource(_fromSource string) error {
    r._fromSource = _fromSource
    r.Set("from_source", _fromSource)
    return nil
}

// FromSource Getter
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetFromSource() string {
    return r._fromSource
}
