package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest
下载飞猪国际城市结果 API请求
taobao.xhotel.city.coordinates.batch.download

给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程 */
type TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest struct {
	model.Params
	// 上传的经纬度批次号
	_batchId int64
}

// New
