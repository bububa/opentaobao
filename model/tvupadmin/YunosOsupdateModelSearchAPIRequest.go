package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdatemodelsearchAPIRequest 机型检索 API请求
// yunos.osupdate.model.search
//
// 机型检索
type YunososupdatemodelsearchAPIRequest struct {
	model.Params
	// 关键词
	_name string
	// 应用ID
	_appId int64
}

// NewYunososupdatemodelsearchRequest 初始化YunososupdatemodelsearchAPIRequest对象
func NewYunososupdatemodelsearchRequest() *YunososupdatemodelsearchAPIRequest {
	return &YunososupdatemodelsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdatemodelsearchAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.model.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdatemodelsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdatemodelsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 关键词
func (r *YunososupdatemodelsearchAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunososupdatemodelsearchAPIRequest) GetName() string {
	return r._name
}

// SetAppId is AppId Setter
// 应用ID
func (r *YunososupdatemodelsearchAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r YunososupdatemodelsearchAPIRequest) GetAppId() int64 {
	return r._appId
}
