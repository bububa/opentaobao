package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMjOcGetproductbyscancode
POS商品查询接口
alibaba.mj.oc.getproductbyscancode

此API用于在银泰商场中，POS端扫码获取商品信息 */
func AlibabaMjOcGetproductbyscancode(clt *core.SDKClient, req *mos.AlibabaMjOcGetproductbyscancodeAPIRequest, session string) (*mos.AlibabaMjOcGetproductbyscancodeAPIResponse, error) {
	var resp mos.AlibabaMjOcGetproductbyscancodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
