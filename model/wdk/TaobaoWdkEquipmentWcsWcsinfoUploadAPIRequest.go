package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest
悬挂链业务信息上传 API请求
taobao.wdk.equipment.wcs.wcsinfo.upload

五道口仓库悬挂链信息上传 */
type TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest struct {
	model.Params
	// 上传信息
	_param0 string
}

// New
