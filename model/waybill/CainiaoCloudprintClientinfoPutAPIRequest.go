package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintclientinfoputAPIRequest 云打印客户端监控信息收集 API请求
// cainiao.cloudprint.clientinfo.put
//
// 云打印客户端监控信息收集
type CainiaocloudprintclientinfoputAPIRequest struct {
	model.Params
	// 客户端上传json数据
	_jsonData string
}

// NewCainiaocloudprintclientinfoputRequest 初始化CainiaocloudprintclientinfoputAPIRequest对象
func NewCainiaocloudprintclientinfoputRequest() *CainiaocloudprintclientinfoputAPIRequest {
	return &CainiaocloudprintclientinfoputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintclientinfoputAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.clientinfo.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintclientinfoputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintclientinfoputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJsonData is JsonData Setter
// 客户端上传json数据
func (r *CainiaocloudprintclientinfoputAPIRequest) SetJsonData(_jsonData string) error {
	r._jsonData = _jsonData
	r.Set("json_data", _jsonData)
	return nil
}

// GetJsonData JsonData Getter
func (r CainiaocloudprintclientinfoputAPIRequest) GetJsonData() string {
	return r._jsonData
}
