package tmallcar

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoWisdomdataOmidRecieveAPIRequest) Reset() {
	r._modelConfig = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoWisdomdataOmidRecieveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.wisdomdata.omid.recieve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoWisdomdataOmidRecieveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoWisdomdataOmidRecieveAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallAliautoWisdomdataOmidRecieveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoWisdomdataOmidRecieveRequest()
	},
}

// GetTmallAliautoWisdomdataOmidRecieveRequest 从 sync.Pool 获取 TmallAliautoWisdomdataOmidRecieveAPIRequest
func GetTmallAliautoWisdomdataOmidRecieveAPIRequest() *TmallAliautoWisdomdataOmidRecieveAPIRequest {
	return poolTmallAliautoWisdomdataOmidRecieveAPIRequest.Get().(*TmallAliautoWisdomdataOmidRecieveAPIRequest)
}

// ReleaseTmallAliautoWisdomdataOmidRecieveAPIRequest 将 TmallAliautoWisdomdataOmidRecieveAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoWisdomdataOmidRecieveAPIRequest(v *TmallAliautoWisdomdataOmidRecieveAPIRequest) {
	v.Reset()
	poolTmallAliautoWisdomdataOmidRecieveAPIRequest.Put(v)
}
