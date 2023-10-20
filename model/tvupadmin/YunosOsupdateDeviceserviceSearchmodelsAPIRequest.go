package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdatedeviceservicesearchmodelsAPIRequest 根据关键词检索设备型号 API请求
// yunos.osupdate.deviceservice.searchmodels
//
// 根据关键词检索设备型号
type YunososupdatedeviceservicesearchmodelsAPIRequest struct {
	model.Params
	// 关键词
	_name string
	// 设备父ID
	_parentId int64
}

// NewYunososupdatedeviceservicesearchmodelsRequest 初始化YunososupdatedeviceservicesearchmodelsAPIRequest对象
func NewYunososupdatedeviceservicesearchmodelsRequest() *YunososupdatedeviceservicesearchmodelsAPIRequest {
	return &YunososupdatedeviceservicesearchmodelsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdatedeviceservicesearchmodelsAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.deviceservice.searchmodels"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdatedeviceservicesearchmodelsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdatedeviceservicesearchmodelsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 关键词
func (r *YunososupdatedeviceservicesearchmodelsAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunososupdatedeviceservicesearchmodelsAPIRequest) GetName() string {
	return r._name
}

// SetParentId is ParentId Setter
// 设备父ID
func (r *YunososupdatedeviceservicesearchmodelsAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r YunososupdatedeviceservicesearchmodelsAPIRequest) GetParentId() int64 {
	return r._parentId
}
