package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautowisdomdataomidrecieveAPIRequest 大搜车车型参配数据接入 API请求
// tmall.aliauto.wisdomdata.omid.recieve
//
// 大搜车车型参配数据接入
type TmallaliautowisdomdataomidrecieveAPIRequest struct {
	model.Params
	// 大搜车车型参配报文
	_modelConfig string
}

// NewTmallaliautowisdomdataomidrecieveRequest 初始化TmallaliautowisdomdataomidrecieveAPIRequest对象
func NewTmallaliautowisdomdataomidrecieveRequest() *TmallaliautowisdomdataomidrecieveAPIRequest {
	return &TmallaliautowisdomdataomidrecieveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautowisdomdataomidrecieveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.wisdomdata.omid.recieve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautowisdomdataomidrecieveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautowisdomdataomidrecieveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModelConfig is ModelConfig Setter
// 大搜车车型参配报文
func (r *TmallaliautowisdomdataomidrecieveAPIRequest) SetModelConfig(_modelConfig string) error {
	r._modelConfig = _modelConfig
	r.Set("model_config", _modelConfig)
	return nil
}

// GetModelConfig ModelConfig Getter
func (r TmallaliautowisdomdataomidrecieveAPIRequest) GetModelConfig() string {
	return r._modelConfig
}
