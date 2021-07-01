package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

/* AlitripRailIrStationGet
国际火车票标准车站查询
alitrip.rail.ir.station.get

国际火车票提供给代理商用于查询标准车站信息，用于代理商对自己的车站与飞猪平台的车站做映射 */
func AlitripRailIrStationGet(clt *core.SDKClient, req *rail.AlitripRailIrStationGetAPIRequest, session string) (*rail.AlitripRailIrStationGetAPIResponse, error) {
	var resp rail.AlitripRailIrStationGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
