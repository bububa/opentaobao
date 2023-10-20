package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshoweditAPIRequest 媒资节目信息修改 API请求
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
type YunostvpubadmincontentshoweditAPIRequest struct {
	model.Params
	// 请求入参，JSON格式
	_data string
}

// NewYunostvpubadmincontentshoweditRequest 初始化YunostvpubadmincontentshoweditAPIRequest对象
func NewYunostvpubadmincontentshoweditRequest() *YunostvpubadmincontentshoweditAPIRequest {
	return &YunostvpubadmincontentshoweditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentshoweditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentshoweditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentshoweditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 请求入参，JSON格式
func (r *YunostvpubadmincontentshoweditAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r YunostvpubadmincontentshoweditAPIRequest) GetData() string {
	return r._data
}
