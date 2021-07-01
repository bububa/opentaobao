package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

/* TaobaoAlitripItFareDelete
【国际机票自有政策】单条删除
taobao.alitrip.it.fare.delete

自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败 */
func TaobaoAlitripItFareDelete(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareDeleteAPIRequest, session string) (*itpolicy.TaobaoAlitripItFareDeleteAPIResponse, error) {
	var resp itpolicy.TaobaoAlitripItFareDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
