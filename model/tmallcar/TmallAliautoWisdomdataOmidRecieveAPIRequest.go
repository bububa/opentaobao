package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoWisdomdataOmidRecieveAPIRequest 大搜车车型参配数据接入 API请求
// tmall.aliauto.wisdomdata.omid.recieve
//
// 大搜车车型参配数据接入
type TmallAliautoWisdomdataOmidRecieveAPIRequest struct {
	model.Params
	// 大搜车车型参配报文
	_modelConfig string
}

// NewTmallAliautoWisdomdataOmidRecieveRequest 初始化TmallAliautoWisdomdataOmidRecieveAPIRequest对象
func NewTmallAliautoWisdomdataOmidRecieveRequest() *TmallAliautoWisdomdataOmidRecieveAPIRequest {
	return &TmallAliautoWisdomdataOmidRecieveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoWisdomdataOmidRecieveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.wisdomdata.omid.recieve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoWisdomdataOmidRecieveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetModelConfig is ModelConfig Setter
// 大搜车车型参配报文
func (r *TmallAliautoWisdomdataOmidRecieveAPIRequest) SetModelConfig(_modelConfig string) error {
	r._modelConfig = _modelConfig
	r.Set("model_config", _modelConfig)
	return nil
}

// GetModelConfig ModelConfig Getter
func (r TmallAliautoWisdomdataOmidRecieveAPIRequest) GetModelConfig() string {
	return r._modelConfig
}
