package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerPointsSync 经纪人积分同步
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
func AlibabaAlihouseExistinghomeBrokerPointsSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
