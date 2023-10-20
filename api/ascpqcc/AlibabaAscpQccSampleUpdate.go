package ascpqcc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpqcc"
)

// AlibabaAscpQccSampleUpdate 品控中心更新样品信息
// alibaba.ascp.qcc.sample.update
//
// 品控中心更新样品信息
func AlibabaAscpQccSampleUpdate(clt *core.SDKClient, req *ascpqcc.AlibabaAscpQccSampleUpdateAPIRequest, session string) (*ascpqcc.AlibabaAscpQccSampleUpdateAPIResponse, error) {
	var resp ascpqcc.AlibabaAscpQccSampleUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
