package autonavi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAutonaviApiTrafficboardImageGetAPIRequest
交通看板-栅格情报获取 API请求
alibaba.autonavi.api.trafficboard.image.get

获取指定情报板ID的二进制数据（图片） */
type AlibabaAutonaviApiTrafficboardImageGetAPIRequest struct {
	model.Params
	// 设备id,  按照userid 的配置，决定是否需要
	_deviceid string
	// 批次,终端批次，按照userid 的配置，决定是否需要
	_batch string
	// 图片 id
	_panelid string
	// 图像尺寸（可选）,默认尺寸为原始大小(960x600) 参数为:width  x   height   (例如:960x600),参数不正确时返回原始大小
	_size string
	// 是否为宽高等比例（可选）,参数值 true（默认）,表示宽高等比例缩放 false:  按请求尺寸缩放
	_whscale string
	// 城市编码
	_adcodes string
}

// New
