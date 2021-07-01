package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintCustomareaUpdateAPIRequest
自定义区内容更新 API请求
cainiao.cloudprint.customarea.update

自定义区内容更新 */
type CainiaoCloudprintCustomareaUpdateAPIRequest struct {
	model.Params
	// 自定义区id（不可修改）
	_customAreaId int64
	// 自定义区名称（可修改）
	_customAreaName string
	// 自定义区内容（可修改）
	_customAreaContent string
}

// New
