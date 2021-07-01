package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

/* AlibabaLstPosOpenAccountCheckissettled
校验当前用户是否入驻了零售通门店接口
alibaba.lst.pos.open.account.checkissettled

校验当前用户是否入驻了零售通门店接口 */
func AlibabaLstPosOpenAccountCheckissettled(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenAccountCheckissettledAPIRequest, session string) (*lstpos.AlibabaLstPosOpenAccountCheckissettledAPIResponse, error) {
	var resp lstpos.AlibabaLstPosOpenAccountCheckissettledAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
