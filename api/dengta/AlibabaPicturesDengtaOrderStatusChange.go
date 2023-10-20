package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// Alibabapicturesdengtaorderstatuschange 天下秀订单状态变更通知
// alibaba.pictures.dengta.order.status.change
//
// 天下秀订单状态变更通知
func Alibabapicturesdengtaorderstatuschange(clt *core.SDKClient, req *dengta.AlibabapicturesdengtaorderstatuschangeAPIRequest, session string) (*dengta.AlibabapicturesdengtaorderstatuschangeAPIResponse, error) {
	var resp dengta.AlibabapicturesdengtaorderstatuschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
