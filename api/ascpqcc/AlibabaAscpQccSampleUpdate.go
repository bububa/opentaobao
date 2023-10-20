package ascpqcc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpqcc"
)

// Alibabaascpqccsampleupdate 品控中心更新样品信息
// alibaba.ascp.qcc.sample.update
//
// 品控中心更新样品信息
func Alibabaascpqccsampleupdate(clt *core.SDKClient, req *ascpqcc.AlibabaascpqccsampleupdateAPIRequest, session string) (*ascpqcc.AlibabaascpqccsampleupdateAPIResponse, error) {
	var resp ascpqcc.AlibabaascpqccsampleupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
