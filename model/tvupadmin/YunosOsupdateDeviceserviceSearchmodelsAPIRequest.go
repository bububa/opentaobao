package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateDeviceserviceSearchmodelsAPIRequest 根据关键词检索设备型号 API请求
// yunos.osupdate.deviceservice.searchmodels
//
// 根据关键词检索设备型号
type YunosOsupdateDeviceserviceSearchmodelsAPIRequest struct {
	model.Params
	// 关键词
	_name string
	// 设备父ID
	_parentId int64
}

// NewYunosOsupdateDeviceserviceSearchmodelsRequest 初始化YunosOsupdateDeviceserviceSearchmodelsAPIRequest对象
func NewYunosOsupdateDeviceserviceSearchmodelsRequest() *YunosOsupdateDeviceserviceSearchmodelsAPIRequest {
	return &YunosOsupdateDeviceserviceSearchmodelsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateDeviceserviceSearchmodelsAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.deviceservice.searchmodels"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateDeviceserviceSearchmodelsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// 关键词
func (r *YunosOsupdateDeviceserviceSearchmodelsAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunosOsupdateDeviceserviceSearchmodelsAPIRequest) GetName() string {
	return r._name
}

// SetParentId is ParentId Setter
// 设备父ID
func (r *YunosOsupdateDeviceserviceSearchmodelsAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r YunosOsupdateDeviceserviceSearchmodelsAPIRequest) GetParentId() int64 {
	return r._parentId
}
