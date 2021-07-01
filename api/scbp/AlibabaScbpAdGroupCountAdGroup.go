package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdGroupCountAdGroup
统计adgroup数量
alibaba.scbp.ad.group.count.ad.group

统计adgroup数量 */
func AlibabaScbpAdGroupCountAdGroup(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupCountAdGroupAPIRequest, session string) (*scbp.AlibabaScbpAdGroupCountAdGroupAPIResponse, error) {
	var resp scbp.AlibabaScbpAdGroupCountAdGroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
