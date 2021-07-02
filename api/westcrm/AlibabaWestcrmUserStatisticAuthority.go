package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmUserStatisticAuthority 获取指定用户是否含有会员权益统计权限
// alibaba.westcrm.user.statistic.authority
//
// 获取指定用户是否含有会员权益统计权限
func AlibabaWestcrmUserStatisticAuthority(clt *core.SDKClient, req *westcrm.AlibabaWestcrmUserStatisticAuthorityAPIRequest, session string) (*westcrm.AlibabaWestcrmUserStatisticAuthorityAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmUserStatisticAuthorityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
