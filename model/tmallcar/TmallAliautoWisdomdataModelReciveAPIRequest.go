package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoWisdomdataModelReciveAPIRequest 第三方车型数据上传 API请求
// tmall.aliauto.wisdomdata.model.recive
//
// 天猫汽车对外提供的汽车车型数据上传接口
type TmallAliautoWisdomdataModelReciveAPIRequest struct {
	model.Params
	// JSON格式车型完整数据
	_modelDetail string
	// 接入的第三方库中的车型唯一id
	_resourceId string
	// 接入的第三方库的名称
	_fromSource string
}

// NewTmallAliautoWisdomdataModelReciveRequest 初始化TmallAliautoWisdomdataModelReciveAPIRequest对象
func NewTmallAliautoWisdomdataModelReciveRequest() *TmallAliautoWisdomdataModelReciveAPIRequest {
	return &TmallAliautoWisdomdataModelReciveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.wisdomdata.model.recive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ModelDetail Setter
// JSON格式车型完整数据
func (r *TmallAliautoWisdomdataModelReciveAPIRequest) SetModelDetail(_modelDetail string) error {
	r._modelDetail = _modelDetail
	r.Set("model_detail", _modelDetail)
	return nil
}

// Get ModelDetail Getter
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetModelDetail() string {
	return r._modelDetail
}

// Set is ResourceId Setter
// 接入的第三方库中的车型唯一id
func (r *TmallAliautoWisdomdataModelReciveAPIRequest) SetResourceId(_resourceId string) error {
	r._resourceId = _resourceId
	r.Set("resource_id", _resourceId)
	return nil
}

// Get ResourceId Getter
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetResourceId() string {
	return r._resourceId
}

// Set is FromSource Setter
// 接入的第三方库的名称
func (r *TmallAliautoWisdomdataModelReciveAPIRequest) SetFromSource(_fromSource string) error {
	r._fromSource = _fromSource
	r.Set("from_source", _fromSource)
	return nil
}

// Get FromSource Getter
func (r TmallAliautoWisdomdataModelReciveAPIRequest) GetFromSource() string {
	return r._fromSource
}
