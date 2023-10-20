package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscconsumerecordsync 外域订单同步本地生活消费记录
// alibaba.alsc.consumerecord.sync
//
// 外部第三方将本地生活app端下单数据同步到本地生活消费记录中心
func Alibabaalscconsumerecordsync(clt *core.SDKClient, req *alsc.AlibabaalscconsumerecordsyncAPIRequest, session string) (*alsc.AlibabaalscconsumerecordsyncAPIResponse, error) {
	var resp alsc.AlibabaalscconsumerecordsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
