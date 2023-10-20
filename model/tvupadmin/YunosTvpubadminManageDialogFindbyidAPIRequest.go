package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagedialogfindbyidAPIRequest 根据id查询全局弹窗 API请求
// yunos.tvpubadmin.manage.dialog.findbyid
//
// 根据id查询全局弹窗
type YunostvpubadminmanagedialogfindbyidAPIRequest struct {
	model.Params
	// 全局弹窗id
	_id int64
}

// NewYunostvpubadminmanagedialogfindbyidRequest 初始化YunostvpubadminmanagedialogfindbyidAPIRequest对象
func NewYunostvpubadminmanagedialogfindbyidRequest() *YunostvpubadminmanagedialogfindbyidAPIRequest {
	return &YunostvpubadminmanagedialogfindbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagedialogfindbyidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.findbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagedialogfindbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagedialogfindbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 全局弹窗id
func (r *YunostvpubadminmanagedialogfindbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadminmanagedialogfindbyidAPIRequest) GetId() int64 {
	return r._id
}
