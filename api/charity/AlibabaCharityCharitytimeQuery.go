package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

/* AlibabaCharityCharitytimeQuery
查询公益3小时公益时汇总
alibaba.charity.charitytime.query

查询公益3小时公益时汇总 */
func AlibabaCharityCharitytimeQuery(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeQueryAPIRequest, session string) (*charity.AlibabaCharityCharitytimeQueryAPIResponse, error) {
	var resp charity.AlibabaCharityCharitytimeQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
