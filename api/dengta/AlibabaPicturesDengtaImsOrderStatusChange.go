package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// Alibabapicturesdengtaimsorderstatuschange 天下秀回传订单执行状态变动
// alibaba.pictures.dengta.ims.order.status.change
//
// 天下秀回传订单执行状态变动
func Alibabapicturesdengtaimsorderstatuschange(clt *core.SDKClient, req *dengta.AlibabapicturesdengtaimsorderstatuschangeAPIRequest, session string) (*dengta.AlibabapicturesdengtaimsorderstatuschangeAPIResponse, error) {
	var resp dengta.AlibabapicturesdengtaimsorderstatuschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
