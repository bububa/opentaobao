package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

/* AlibabaNrsItemRtdataBackflow
RT竞价数据回流
alibaba.nrs.item.rtdata.backflow

回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正 */
func AlibabaNrsItemRtdataBackflow(clt *core.SDKClient, req *dt.AlibabaNrsItemRtdataBackflowAPIRequest, session string) (*dt.AlibabaNrsItemRtdataBackflowAPIResponse, error) {
	var resp dt.AlibabaNrsItemRtdataBackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
