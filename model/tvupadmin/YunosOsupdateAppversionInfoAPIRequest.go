package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateappversioninfoAPIRequest 获取应用升级详情 API请求
// yunos.osupdate.appversion.info
//
// 获取应用升级详情
type YunososupdateappversioninfoAPIRequest struct {
	model.Params
	// 升级任务ID
	_id int64
}

// NewYunososupdateappversioninfoRequest 初始化YunososupdateappversioninfoAPIRequest对象
func NewYunososupdateappversioninfoRequest() *YunososupdateappversioninfoAPIRequest {
	return &YunososupdateappversioninfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateappversioninfoAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateappversioninfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateappversioninfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 升级任务ID
func (r *YunososupdateappversioninfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunososupdateappversioninfoAPIRequest) GetId() int64 {
	return r._id
}
