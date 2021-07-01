package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpEffectAccountList
账户-报表
alibaba.scbp.effect.account.list

账户-报表,支持最近7天，最近30天，以及180天内时间区间。 */
func AlibabaScbpEffectAccountList(clt *core.SDKClient, req *scbp.AlibabaScbpEffectAccountListAPIRequest, session string) (*scbp.AlibabaScbpEffectAccountListAPIResponse, error) {
	var resp scbp.AlibabaScbpEffectAccountListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
