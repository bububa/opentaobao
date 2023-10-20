package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectkanamequeryAPIRequest 查询KA楼盘名称 API请求
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
type AlibabaalihousenewhomeprojectkanamequeryAPIRequest struct {
	model.Params
	// KA楼盘ID
	_kaProjectMid int64
}

// NewAlibabaalihousenewhomeprojectkanamequeryRequest 初始化AlibabaalihousenewhomeprojectkanamequeryAPIRequest对象
func NewAlibabaalihousenewhomeprojectkanamequeryRequest() *AlibabaalihousenewhomeprojectkanamequeryAPIRequest {
	return &AlibabaalihousenewhomeprojectkanamequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectkanamequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.kaname.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectkanamequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectkanamequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKaProjectMid is KaProjectMid Setter
// KA楼盘ID
func (r *AlibabaalihousenewhomeprojectkanamequeryAPIRequest) SetKaProjectMid(_kaProjectMid int64) error {
	r._kaProjectMid = _kaProjectMid
	r.Set("ka_project_mid", _kaProjectMid)
	return nil
}

// GetKaProjectMid KaProjectMid Getter
func (r AlibabaalihousenewhomeprojectkanamequeryAPIRequest) GetKaProjectMid() int64 {
	return r._kaProjectMid
}
