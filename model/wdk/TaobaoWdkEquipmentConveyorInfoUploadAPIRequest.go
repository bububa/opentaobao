package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentConveyorInfoUploadAPIRequest
五道口仓库悬挂链信息上报 API请求
taobao.wdk.equipment.conveyor.info.upload

五道口仓库悬挂链信息上传 */
type TaobaoWdkEquipmentConveyorInfoUploadAPIRequest struct {
	model.Params
	// 上传信息
	_param0 string
}

// New
