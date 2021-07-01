package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest
上传小件员GPS位置信息 API请求
cainiao.guoguo.cp.nborderfrontr.uploadcoordinate

上传小件员GPS位置信息 */
type CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest struct {
	model.Params
	// 小件员所在公司编号
	_cpCode string
	// 小件员员工编号
	_cpUserId string
	// 上报时间，格式：yyyy-MM-dd HH:mm:ss
	_timeStamp string
	// 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
	_source int64
	// 经度
	_lng string
	// 纬度
	_lat string
	// 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
	_gpsType string
}

// New
