package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiMobileFigureimageGetAPIRequest
获取手机banner图 API请求
alibaba.xiami.api.mobile.figureimage.get

获取手机banner图 */
type AlibabaXiamiApiMobileFigureimageGetAPIRequest struct {
	model.Params
	// 分页限制
	_limit int64
	// 类型
	_type string
	// 客户端版本
	_av string
	// 设备类型
	_deviceType string
	// 设备ID
	_deviceId string
}

// New
