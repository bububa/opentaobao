package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest 更新楼盘商品信息 API请求
// alibaba.alihouse.newhome.project.update.item.info
//
// 更新楼盘商品信息
type AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest struct {
	model.Params
	// 实体类
	_projectUpdateItemInfoDto *ProjectUpdateItemInfoDto
}

// NewAlibabaalihousenewhomeprojectupdateiteminfoRequest 初始化AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest对象
func NewAlibabaalihousenewhomeprojectupdateiteminfoRequest() *AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest {
	return &AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.update.item.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectUpdateItemInfoDto is ProjectUpdateItemInfoDto Setter
// 实体类
func (r *AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest) SetProjectUpdateItemInfoDto(_projectUpdateItemInfoDto *ProjectUpdateItemInfoDto) error {
	r._projectUpdateItemInfoDto = _projectUpdateItemInfoDto
	r.Set("project_update_item_info_dto", _projectUpdateItemInfoDto)
	return nil
}

// GetProjectUpdateItemInfoDto ProjectUpdateItemInfoDto Getter
func (r AlibabaalihousenewhomeprojectupdateiteminfoAPIRequest) GetProjectUpdateItemInfoDto() *ProjectUpdateItemInfoDto {
	return r._projectUpdateItemInfoDto
}
