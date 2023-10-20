package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminadmottauditAPIRequest 优酷OTT广告素材审核 API请求
// yunos.tvpubadmin.adm.ott.audit
//
// 审核优酷OTT端广告素材
type YunostvpubadminadmottauditAPIRequest struct {
	model.Params
	// 广告审核内容，json格式
	_data string
}

// NewYunostvpubadminadmottauditRequest 初始化YunostvpubadminadmottauditAPIRequest对象
func NewYunostvpubadminadmottauditRequest() *YunostvpubadminadmottauditAPIRequest {
	return &YunostvpubadminadmottauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminadmottauditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.adm.ott.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminadmottauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminadmottauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 广告审核内容，json格式
func (r *YunostvpubadminadmottauditAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r YunostvpubadminadmottauditAPIRequest) GetData() string {
	return r._data
}
