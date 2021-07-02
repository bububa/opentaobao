package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowEditAPIRequest 媒资节目信息修改 API请求
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
type YunosTvpubadminContentShowEditAPIRequest struct {
	model.Params
	// 请求入参，JSON格式
	_data string
}

// NewYunosTvpubadminContentShowEditRequest 初始化YunosTvpubadminContentShowEditAPIRequest对象
func NewYunosTvpubadminContentShowEditRequest() *YunosTvpubadminContentShowEditAPIRequest {
	return &YunosTvpubadminContentShowEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowEditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Data Setter
// 请求入参，JSON格式
func (r *YunosTvpubadminContentShowEditAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r YunosTvpubadminContentShowEditAPIRequest) GetData() string {
	return r._data
}
