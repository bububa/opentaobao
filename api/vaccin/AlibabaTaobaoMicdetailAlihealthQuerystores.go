package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

/* AlibabaTaobaoMicdetailAlihealthQuerystores
疫苗预约门店列表查询
alibaba.taobao.micdetail.alihealth.querystores

微信小程序疫苗预约门店列表查询 */
func AlibabaTaobaoMicdetailAlihealthQuerystores(clt *core.SDKClient, req *vaccin.AlibabaTaobaoMicdetailAlihealthQuerystoresAPIRequest, session string) (*vaccin.AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponse, error) {
	var resp vaccin.AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
