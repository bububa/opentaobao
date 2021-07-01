package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

/* AlibabaWestcrmUserConsumerGet
获取指定用户的消费总额
alibaba.westcrm.user.consumer.get

获取指定用户的消费总额 */
func AlibabaWestcrmUserConsumerGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmUserConsumerGetAPIRequest, session string) (*westcrm.AlibabaWestcrmUserConsumerGetAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmUserConsumerGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
