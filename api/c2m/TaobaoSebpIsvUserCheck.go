package c2m

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/c2m"
)

/* TaobaoSebpIsvUserCheck
淘小铺账户实名校验接口
taobao.sebp.isv.user.check

校验淘小铺账户和身份信息匹配成功 */
func TaobaoSebpIsvUserCheck(clt *core.SDKClient, req *c2m.TaobaoSebpIsvUserCheckAPIRequest, session string) (*c2m.TaobaoSebpIsvUserCheckAPIResponse, error) {
	var resp c2m.TaobaoSebpIsvUserCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
