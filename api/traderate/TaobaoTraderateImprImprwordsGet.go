package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

/* TaobaoTraderateImprImprwordsGet
评价大家印象印象短语接口
taobao.traderate.impr.imprwords.get

根据淘宝后台类目的一级类目和叶子类目 */
func TaobaoTraderateImprImprwordsGet(clt *core.SDKClient, req *traderate.TaobaoTraderateImprImprwordsGetAPIRequest, session string) (*traderate.TaobaoTraderateImprImprwordsGetAPIResponse, error) {
	var resp traderate.TaobaoTraderateImprImprwordsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
