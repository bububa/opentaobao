package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// AlibabaNrsItemRtdataBackflow RT竞价数据回流
// alibaba.nrs.item.rtdata.backflow
//
// 回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
func AlibabaNrsItemRtdataBackflow(clt *core.SDKClient, req *dt.AlibabaNrsItemRtdataBackflowAPIRequest, resp *dt.AlibabaNrsItemRtdataBackflowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
