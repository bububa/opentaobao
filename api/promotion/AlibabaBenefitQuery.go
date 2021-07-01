package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* AlibabaBenefitQuery
奖池奖品查询列表
alibaba.benefit.query

功能：奖池奖品查询列表
业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池提供的奖品返回给
小程。
安全保障：为保证数据不会越权，需要卖家授，并且验证系统参数appKey。只有通过授权的，并且
appkey验证通过的，才会查数据 并透出，否则直接失败。
因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
是否一致，所以可以保证参数appName的合法性，没有越权。 */
func AlibabaBenefitQuery(clt *core.SDKClient, req *promotion.AlibabaBenefitQueryAPIRequest, session string) (*promotion.AlibabaBenefitQueryAPIResponse, error) {
	var resp promotion.AlibabaBenefitQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
